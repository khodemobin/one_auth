package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/khodemobin/pilo/auth/app"
	"github.com/khodemobin/pilo/auth/internal/model"
	"github.com/khodemobin/pilo/auth/pkg/helper"
	"gorm.io/gorm"
)

type userRepo struct{}

func NewUserRepo() UserRepository {
	return &userRepo{}
}

func (u *userRepo) Find(ctx context.Context, column string, value string) (*model.User, error) {
	var user model.User

	cache, err := u.getFromCache(column, value)
	if err != nil {
		return cache, err
	}

	err = app.DB().Where(column+" = ? ", value).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, app.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	err = u.saveToCache(&user, column, value)

	return &user, err
}

func (u *userRepo) FindActive(ctx context.Context, column string, value string) (*model.User, error) {
	var user *model.User

	user, err := u.getFromCache(column, value)
	if err != nil && user.IsActive {
		return user, err
	}

	err = app.DB().Where(column+" = ? ", value).Where("is_active", 1).First(user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, app.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	err = u.saveToCache(user, column, value)

	return user, err
}

func (u *userRepo) UpdateLastSeen(ctx context.Context, user *model.User) error {
	now := time.Now()
	user.LastSignInAt = &now
	_, err := u.Update(ctx, user)
	return err
}

func (u *userRepo) Update(ctx context.Context, user *model.User) (*model.User, error) {
	err := app.DB().Save(user).Error
	if err != nil {
		return nil, err
	}
	u.resetCache(user)
	return user, nil
}

func (*userRepo) Create(ctx context.Context, user *model.User) (*model.User, error) {
	err := app.DB().Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (*userRepo) ExistsByID(ctx context.Context, id uint) (bool, error) {
	var count int64
	err := app.DB().Table("users").Where("id = ?", id).Count(&count).Error

	return count > 0, err
}

func (*userRepo) ExistsByPhone(ctx context.Context, phone string) (bool, error) {
	var count int64
	err := app.DB().Table("users").Where("phone = ?", phone).Count(&count).Error

	return count > 0, err
}

func (*userRepo) ExistsByUUID(ctx context.Context, uuid string) (bool, error) {
	var count int64
	err := app.DB().Table("users").Where("uuid = ?", uuid).Count(&count).Error

	return count > 0, err
}

func (*userRepo) resetCache(user *model.User) {
	app.Cache().Delete(fmt.Sprintf("user_by_uuid_%s", user.UUID))
	app.Cache().Delete(fmt.Sprintf("user_by_id_%d", user.ID))
	app.Cache().Delete(fmt.Sprintf("user_by_uuid_%s", *user.Phone))
}

func (*userRepo) getFromCache(column, value string) (*model.User, error) {
	var user model.User
	cache, err := app.Cache().Get(fmt.Sprintf("user_by_%s_%s", column, value), nil)
	if err == nil && cache != nil {
		err = json.Unmarshal([]byte(*cache), &user)
		return &user, err
	}
	return nil, err
}

func (u *userRepo) saveToCache(user *model.User, column string, value string) error {
	json, err := helper.ToJson(user)
	if err != nil {
		return err
	}

	return app.Cache().Set(fmt.Sprintf("user_by_%s_%s", column, value), json, 0)
}
