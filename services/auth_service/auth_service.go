package auth_service

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"hkn-be/config"
	"hkn-be/constants"
	"hkn-be/infras/jwt_infra"
	"hkn-be/infras/mail"
	"hkn-be/models"
	"hkn-be/objects"
	"hkn-be/repositories"
	"hkn-be/utils"
	"time"
)

type authService struct {
	*repositories.RepositoryCtx
	server config.ServerConfig
	jwt    jwt_infra.JwtInterface
	rd     *redis.Client
	mail   mail.MailServiceInterface
}

func (a authService) Register(ctx context.Context, userData objects.User) error {
	//todo:: temporary set the user tole by default as student, but should be change later if flow is in final version
	//begin of set default user role as student
	roleData, err := a.RoleRepo.GetRoleByCode(ctx, "student")
	if err != nil {
		return err
	}

	userData.Roles = []objects.Role{
		{
			Id:          roleData.Id,
			Code:        roleData.Code,
			Name:        roleData.Name,
			Description: roleData.Description,
			CreatedAt:   roleData.CreatedAt,
			UpdatedAt:   roleData.UpdatedAt,
			DeletedAt:   roleData.DeletedAt,
		},
	}
	//end of set default user role as student

	if utils.ValidateEmail(userData.Email) == false {
		return constants.ErrInvalidEmailAddress
	}
	hashedPassword, err := utils.HashPassword(userData.Password)
	if err != nil {
		return err
	}

	tx := a.DB.Begin()
	//create user data
	createdUserData, err := a.UserRepo.CreateUser(
		ctx, models.User{
			Email:      userData.Email,
			Password:   hashedPassword,
			Name:       userData.Name,
			IsVerified: false,
		},
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	//create user role_repo after insert user data
	if len(userData.Roles) > 0 {
		for _, role := range userData.Roles {
			_, err := a.UserRoleRepo.CreateUserRole(
				ctx, models.UserRole{
					UserId: createdUserData.Id,
					RoleId: role.Id,
				},
			)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	err = a.getVerificationCode(ctx, createdUserData)

	tx.Commit()
	return nil
}

func (a authService) RequestVerificationCode(ctx context.Context, email string) error {
	userData, err := a.UserRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}

	if userData.IsVerified {
		return errors.New("user is already verified")
	}

	verificationCode, err := utils.RetrieveVerificationCode(a.rd, ctx, userData.Id)
	if err == nil && len(verificationCode) > 0 {
		return errors.New("you still have active verification code")
	}

	err = a.getVerificationCode(ctx, userData)
	if err != nil {
		return err
	}
	return nil
}

func (a authService) getVerificationCode(ctx context.Context, userData *models.User) error {
	verificationCode, err := utils.GenerateVerificationCode()
	if err != nil {
		return err
	}

	err = utils.StoreVerificationCode(a.rd, ctx, userData.Id, verificationCode, 10*time.Minute)
	if err != nil {
		return err
	}

	verificationLink := utils.GenerateVerificationURL(fmt.Sprintf("%s%s/verify-registration", a.server.BaseUrl, a.server.Port), userData.Id, verificationCode)
	emailData := struct {
		Username         string
		VerificationLink string
		VerificationCode string
	}{
		Username:         userData.Name,
		VerificationLink: verificationLink,
		VerificationCode: verificationCode,
	}
	emailBody, err := utils.GenerateEmailBody(constants.RegistrationVerificationMailTemplate, emailData)

	err = a.mail.SendEmail(
		[]string{userData.Email},
		"Registration on HKN",
		emailBody,
	)
	if err != nil {
		return err
	}
	return nil
}

func (a authService) VerifyUserRegistration(ctx context.Context, request objects.VerifyRegistrationRequest) error {
	userData, err := a.UserRepo.GetUserById(ctx, request.UserId)
	if err != nil {
		return err
	}

	if userData.IsVerified {
		return errors.New("user already verified")
	}

	isValid, err := utils.VerifyVerificationCode(a.rd, ctx, request.UserId, request.VerificationCode)
	if err != nil {
		return err
	}
	if !isValid {
		return errors.New("invalid verification code")
	}

	err = a.UserRepo.UpdateUser(
		ctx, models.User{
			Id:         userData.Id,
			IsVerified: true,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (a authService) Login(ctx context.Context, loginRequestData objects.LoginRequest) (*objects.LoginResponse, error) {
	userData, err := a.UserRepo.GetUserByEmail(ctx, loginRequestData.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		} else {
			return nil, err
		}
	}

	if !utils.CheckPassword(userData.Password, loginRequestData.Password) {
		return nil, errors.New("invalid password")
	}

	tmpUserRoles, err := a.UserRoleRepo.GetUserRoleListByUserId(ctx, userData.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user role not found")
		} else {
			return nil, err
		}
	}

	var userRoles []objects.Role

	for _, userRole := range *tmpUserRoles {
		tmpRole, err := a.RoleRepo.GetRoleById(ctx, userRole.RoleId)
		if err != nil {
			return nil, err
		}
		userRoles = append(
			userRoles, objects.Role{
				Id:          tmpRole.Id,
				Code:        tmpRole.Code,
				Name:        tmpRole.Name,
				Description: tmpRole.Description,
				CreatedAt:   tmpRole.CreatedAt,
				UpdatedAt:   tmpRole.UpdatedAt,
				DeletedAt:   tmpRole.DeletedAt,
			},
		)
	}

	userObj := objects.User{
		Id:         userData.Id,
		Email:      userData.Email,
		Password:   "*****",
		Name:       userData.Name,
		IsVerified: userData.IsVerified,
		Roles:      userRoles,
	}

	claims, err := a.jwt.GenerateJwtToken(ctx, userObj)
	if err != nil {
		return nil, err
	}

	loginResponse := objects.LoginResponse{
		UserData: userObj,
		JwtToken: *claims,
	}
	return &loginResponse, nil
}

func (a authService) Logout(ctx context.Context, id string) error {
	err := a.jwt.DeleteTokenFromRedis(ctx, id, jwt_infra.AuthKey)
	if err != nil {
		return err
	}
	return nil
}
