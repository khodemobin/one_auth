package repository

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/khodemobin/pilo/auth/internal/domain"
	"github.com/khodemobin/pilo/auth/pkg/cache"
	"github.com/khodemobin/pilo/auth/pkg/encrypt"
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

func (c *confirmCode) CreateConfirmCode(phone string) error {
	code, originalCode, err := encrypt.GenerateConfirmCode(phone)
	if err != nil {
		return err
	}
	log.Println(originalCode)

	json, err := helper.ToJson(code)
	if err != nil {
		return err
	}

	return c.cache.Set(fmt.Sprintf("user_confirm_code_%s", phone), json, code.ExpiresIn)
}

func (c *confirmCode) FindConfirmCode(phone string) (*domain.ConfirmCode, error) {
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

func (c *confirmCode) DeleteConfirmCode(phone string) error {
	return c.cache.Delete(fmt.Sprintf("user_confirm_code_%s", phone))
}
