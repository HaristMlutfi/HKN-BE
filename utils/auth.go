package utils

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"math/big"
	"net/url"
	"regexp"
	"time"
)

var (
	emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

// ValidateEmail return true if email address has a valid format
func ValidateEmail(email string) bool {
	return emailRegex.MatchString(email)
}

// HashPassword hashes a plain text password using bcrypt
func HashPassword(password string) (string, error) {
	if err := validatePasswordStrength(password); err != nil {
		return "", err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPassword checks if the provided password matches the hashed password
func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// validatePasswordStrength validates the strength of a password
func validatePasswordStrength(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#~$%^&*()+|_.,<>?]`).MatchString(password)

	if !hasUpper {
		return errors.New("password must contain at least one uppercase letter")
	}
	if !hasLower {
		return errors.New("password must contain at least one lowercase letter")
	}
	if !hasNumber {
		return errors.New("password must contain at least one number")
	}
	if !hasSpecial {
		return errors.New("password must contain at least one special character")
	}

	return nil
}

// GenerateVerificationCode generates a random verification code of a given length
func GenerateVerificationCode() (string, error) {
	max := big.NewInt(1000000) // 1 million
	code, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%06d", code), nil
}

// StoreVerificationCode stores a verification code in Redis with an expiration time
func StoreVerificationCode(r *redis.Client, ctx context.Context, userID, code string, expiration time.Duration) error {
	key := fmt.Sprintf("verification_code:%s", userID)
	err := r.Set(ctx, key, code, expiration).Err()
	if err != nil {
		return fmt.Errorf("failed to store verification code: %w", err)
	}
	return nil
}

// RetrieveVerificationCode retrieves the verification code from Redis
func RetrieveVerificationCode(r *redis.Client, ctx context.Context, userID string) (string, error) {
	key := fmt.Sprintf("verification_code:%s", userID)
	code, err := r.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return "", errors.New("verification code not found or expired")
	} else if err != nil {
		return "", fmt.Errorf("failed to retrieve verification code: %w", err)
	}
	return code, nil
}

// VerifyCode checks if the provided code matches the stored verification code
func VerifyVerificationCode(r *redis.Client, ctx context.Context, userID, providedCode string) (bool, error) {
	storedCode, err := RetrieveVerificationCode(r, ctx, userID)
	if err != nil {
		return false, err
	}
	return storedCode == providedCode, nil
}

// GenerateVerificationURL creates a URL with embedded verification parameters
func GenerateVerificationURL(baseURL, userID, verificationCode string) string {
	params := url.Values{}
	params.Add("user_id", userID)
	params.Add("verification_code", verificationCode)
	return fmt.Sprintf("%s?%s", baseURL, params.Encode())
}
