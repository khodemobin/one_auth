package repository

import (
	"encoding/json"
	"fmt"

	"github.com/khodemobin/pilo/auth/internal/domain"
	"github.com/khodemobin/pilo/auth/pkg/cache"
	"github.com/khodemobin/pilo/auth/pkg/helper"
)

type confirmCode struct {
	cache cache.Cache
}

func NewConfirmCodeRepo(cache cache.Cache) domain.ConfirmCodeRepository {
	return &confirmCode{
		cache,
	}
}

func (c confirmCode) Store(phone string, confirmCode *domain.ConfirmCode) error {
	json, err := helper.ToJson(confirmCode)
	if err != nil {
		return err
	}

	return c.cache.Set(fmt.Sprintf("user_confirm_code_%s", phone), json, confirmCode.ExpiresIn)
}

func (c confirmCode) Find(phone string) (*domain.ConfirmCode, error) {
	result, err := c.cache.Get(fmt.Sprintf("user_confirm_code_%s", phone), nil)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	var confirm domain.ConfirmCode
	err = json.Unmarshal([]byte(*result), &confirm)

	return &confirm, err
}
