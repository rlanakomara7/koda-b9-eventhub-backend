package services

import (
	"errors"

	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/models"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo *repositories.UserRepository
}

func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{
		UserRepo: userRepo,
	}
}

// Register
func (s *AuthService) Register(user *models.User) error {
	// cek email
	existingUser, err := s.UserRepo.FindByEmail(user.Email)

	if err != nil {
		return err
	}

	if existingUser != nil {
		return errors.New("email already registered")
	}

	//hash password

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	//default role attendee
	user.RoleID = 1

	//default status
	user.Status = "active"

	//save database
	return s.UserRepo.Creatre(user)
}
