package objects

type LoginRequest struct {
	Email    string
	Password string
}

type RegistrationRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Name     string `json:"name" validate:"required"`
}

type RequestVerificationCodeRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type VerifyRegistrationRequest struct {
	UserId           string `json:"user_id" validate:"required"`
	VerificationCode string `json:"verification_code" validate:"required,min=6"`
}

type LoginResponse struct {
	UserData User
	JwtToken string
}
