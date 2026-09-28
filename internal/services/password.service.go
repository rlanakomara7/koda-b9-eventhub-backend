package services

import (
	"errors"

	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/repositories"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/utils"
)

type PasswordService struct {
	UserRepo     *repositories.UserRepository
	PasswordRepo *repositories.PasswordRepository
}

func NewPasswordService(userRepo *repositories.UserRepository, passwordRepo *repositories.PasswordRepository) *PasswordService {

	return &PasswordService{
		UserRepo:     userRepo,
		PasswordRepo: passwordRepo,
	}
}

// forgot password
func (s *PasswordService) ForgotPassword(email string) (string, error) {

	user, err := s.UserRepo.FindByEmail(email)

	if err != nil {
		return "", err
	}

	if user == nil {
		return "", errors.New("email not found")
	}

	token := utils.GenerateResetToken()

	err = s.PasswordRepo.CreateResetToken(
		user.UserID,
		token,
	)

	if err != nil {
		return "", err
	}

	return token, nil
}

// reset password
func (s *PasswordService) ResetPassword(email string, token string, newPassword string) error {

	user, err := s.UserRepo.FindByEmail(email)

	if err != nil || user == nil {
		return errors.New("user not found")
	}

	userID, err := s.PasswordRepo.FindResetToken(token)

	if err != nil {
		return errors.New("invalid reset token")
	}

	if userID != user.UserID {
		return errors.New("token mismatch")
	}

	hash, err := utils.HashPassword(
		newPassword,
	)

	if err != nil {
		return err
	}

	return s.UserRepo.UpdatePassword(
		user.UserID,
		hash,
	)

}
