package encrypt

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/khodemobin/pilo/auth/app"
	"github.com/khodemobin/pilo/auth/internal/model"

	"github.com/dgrijalva/jwt-go"
)

func GenerateAccessToken(user *model.User) (string, error) {
	secret := app.Config().App.JwtSecret
	ttl, err := strconv.Atoi(app.Config().App.JwtTTL)
	if err != nil {
		return "", err
	}
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * time.Duration(ttl)).Unix(),
		Subject:   user.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func SecureToken() (string, error) {
	b := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}
	return removePadding(base64.URLEncoding.EncodeToString(b)), nil
}

func ParseJWTClaims(bearer string) (string, error) {
	secret := app.Config().App.JwtSecret

	p := jwt.Parser{ValidMethods: []string{jwt.SigningMethodHS256.Name}}
	c, err := p.ParseWithClaims(bearer, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	claims := c.Claims.(*jwt.StandardClaims)

	return claims.Subject, err
}

func removePadding(token string) string {
	return strings.TrimRight(token, "=")
}
