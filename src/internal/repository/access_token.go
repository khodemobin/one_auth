package repository

import (
	"fmt"
	"strconv"
	"time"

	"github.com/khodemobin/pilo/auth/app"
)

type at struct{}

func NewAccessTokenRepo() AccessTokenRepository {
	return &at{}
}

func (*at) AddToBlacklist(token string) error {
	ttl, err := strconv.Atoi(app.Config().App.JwtTTL)
	if err != nil {
		return err
	}

	return app.Cache().Set(fmt.Sprintf("black_list_%s", token), true, time.Second*time.Duration(ttl))
}

func (*at) ExistsInBlackList(token string) (bool, error) {
	value, err := app.Cache().Get(fmt.Sprintf("black_list_%s", token), nil)
	if err != nil {
		return false, err
	}

	return value != nil, nil
}
