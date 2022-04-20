package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/khodemobin/pilo/auth/app"
	"github.com/khodemobin/pilo/auth/internal/model"
	"github.com/khodemobin/pilo/auth/pkg/helper"
	"gorm.io/gorm"
)

type userRepo struct{}

func NewUserRepo() UserRepository {
	return &userRepo{}
}

func (userRepo) FindUserByUUID(ctx context.Context, uuid string, status int) (*model.User, error) {
	cache, err := app.Cache().Remember(fmt.Sprintf("user_by_uuid_%s", uuid), func() (*string, error) {
		var user *model.User

		findQ := &model.User{UUID: uuid}
		if status != -1 {
			findQ.Status = status
		}

		err := app.DB().Where(findQ).First(&user).Error
		err = checkError(err)
		if err != nil {
			return nil, err
		}

		json, err := helper.ToJson(user)
		if err != nil {
			return nil, err
		}

		return &json, nil
	})
	if err != nil {
		return nil, err
	}

	var user model.User
	err = json.Unmarshal([]byte(*cache), &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (userRepo) FindUserByID(ctx context.Context, id uint, status int) (*model.User, error) {
	cache, err := app.Cache().Remember(fmt.Sprintf("user_by_id_%d", id), func() (*string, error) {
		var user *model.User

		findQ := &model.User{ID: id}
		if status != -1 {
			findQ.Status = status
		}
		err := app.DB().Where(findQ).First(&user).Error
		err = checkError(err)
		if err != nil {
			return nil, err
		}

		json, err := helper.ToJson(user)
		if err != nil {
			return nil, err
		}

		return &json, nil
	})
	if err != nil {
		return nil, err
	}

	var user model.User
	err = json.Unmarshal([]byte(*cache), &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (userRepo) FindUserByPhone(ctx context.Context, phone string, status int) (*model.User, error) {
	cache, err := app.Cache().Remember(fmt.Sprintf("user_by_phone_%s", phone), func() (*string, error) {
		var user *model.User

		findQ := &model.User{Phone: phone}
		if status != -1 {
			findQ.Status = status
		}
		err := app.DB().Where(findQ).First(&user).Error
		err = checkError(err)
		if err != nil {
			return nil, err
		}

		json, err := helper.ToJson(user)
		if err != nil {
			return nil, err
		}

		return &json, nil
	})
	if err != nil {
		return nil, err
	}

	var user model.User
	err = json.Unmarshal([]byte(*cache), &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (userRepo) UpdateUserLastSeen(ctx context.Context, user *model.User) error {
	now := time.Now()
	user.LastSignInAt = &now
	err := app.DB().Save(user).Error
	resetCache(user)
	return err
}

func (userRepo) CreateOrUpdateUser(ctx context.Context, user *model.User) error {
	newUser := &model.User{}

	err := app.DB().Where(model.User{Phone: user.Phone}).First(&newUser).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if err != nil {
		user.UUID = uuid.New().String()
		err = app.DB().Create(&user).Error
	} else {
		err = app.DB().Model(&newUser).Updates(user).Error
		resetCache(user)
	}

	return err
}

func (*userRepo) UpdatePassword(ctx context.Context, user *model.User) error {
	panic("unimplemented")
}

func resetCache(user *model.User) {
	app.Cache().Delete(fmt.Sprintf("user_by_uuid_%s", user.UUID))
	app.Cache().Delete(fmt.Sprintf("user_by_id_%d", user.ID))
	app.Cache().Delete(fmt.Sprintf("user_by_uuid_%s", user.Phone))
}
