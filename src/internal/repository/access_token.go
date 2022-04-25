package repository

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/khodemobin/pilo/auth/app"
)

type at struct{}

func NewAccessTokenRepo() AccessTokenRepository {
	return &at{}
}

// AddToBlacklist implements AccessTokenRepository
func (*at) AddToBlacklist(token string) error {
	ttl, err := strconv.Atoi(app.Config().App.JwtTTL)
	if err != nil {
		return err
	}

	return app.Cache().Set(fmt.Sprintf("black_list_%s", token), true, time.Second*time.Duration(ttl))
}

// ExistsInBlackList implements AccessTokenRepository
func (*at) ExistsInBlackList(token string) (bool, error) {
	_, err := app.Cache().Get(fmt.Sprintf("black_list_%s", token), nil)
	if errors.Is(err, redis.Nil) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}
