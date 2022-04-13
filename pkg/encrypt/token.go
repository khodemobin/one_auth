package encrypt

import (
	"time"

	"github.com/khodemobin/pilo/auth/internal/config"
	"github.com/khodemobin/pilo/auth/internal/model"

	"github.com/dgrijalva/jwt-go"
)

func GenerateAccessToken(user *model.User, expiresIn time.Duration) (string, error) {
	secret := config.GetConfig().App.JwtSecret
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(expiresIn).Unix(),
		Subject:   user.UUID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ParseJWTClaims(bearer string) error {
	secret := config.GetConfig().App.JwtSecret

	p := jwt.Parser{ValidMethods: []string{jwt.SigningMethodHS256.Name}}
	_, err := p.ParseWithClaims(bearer, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	return err
}
