package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/config"
)

func GenerateToken(userID uint, email string, roleID uint) (string, error) {

	claims := jwt.MapClaims{

		"user_id": userID,
		"email":   email,
		"role_id": roleID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(
		[]byte(config.GetJWTSecret()),
	)
}
