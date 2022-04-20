package repository

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/khodemobin/pilo/auth/app"
	"github.com/khodemobin/pilo/auth/internal/model"
	"github.com/khodemobin/pilo/auth/pkg/encrypt"
	"github.com/khodemobin/pilo/auth/pkg/helper"
)

type confirmCode struct {
	app *app.AppContainer
}

func NewConfirmCodeRepo() ConfirmCodeRepository {
	app := app.App()
	return &confirmCode{
		app: app,
	}
}

func (c *confirmCode) Create(phone string) error {
	code, originalCode, err := encrypt.GenerateConfirmCode(phone)
	if err != nil {
		return err
	}
	log.Println(originalCode)

	json, err := helper.ToJson(code)
	if err != nil {
		return err
	}

	return c.app.Cache.Set(fmt.Sprintf("user_confirm_code_%s", phone), json, code.ExpiresIn)
}

func (c *confirmCode) Find(phone string) (*model.ConfirmCode, error) {
	result, err := c.app.Cache.Get(fmt.Sprintf("user_confirm_code_%s", phone), nil)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	var confirm model.ConfirmCode
	err = json.Unmarshal([]byte(*result), &confirm)

	return &confirm, err
}

func (c *confirmCode) Delete(phone string) error {
	return c.app.Cache.Delete(fmt.Sprintf("user_confirm_code_%s", phone))
}
