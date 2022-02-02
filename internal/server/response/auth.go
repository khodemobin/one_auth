package response

import "github.com/khodemobin/pilo/auth/internal/domain"

type AuthResource struct {
	Token string `json:"token"`
}

func NewAuthResource(u *domain.Auth) *AuthResource {
	return &AuthResource{
		Token: u.Token,
	}
}
