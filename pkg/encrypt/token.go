package encrypt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/khodemobin/pilo/auth/internal/domain"
)

func GenerateAccessToken(user *domain.User, expiresIn time.Duration, secret string) (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(expiresIn).Unix(),
		Subject:   string(rune(user.ID)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}