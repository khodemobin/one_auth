package repository

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/khodemobin/pilo/auth/app"
	"github.com/khodemobin/pilo/auth/internal/model"
	"gorm.io/gorm"
)

type userRepo struct{}

func NewUserRepo() UserRepository {
	return &userRepo{}
}

func (userRepo) FindUserById(ctx context.Context, id int, status int) (*model.User, error) {
	var user *model.User

	findQ := &model.User{ID: uint(id)}
	if status != -1 {
		findQ.Status = status
	}

	err := app.DB().Where(findQ).First(&user).Error
	if err == nil || errors.Is(err, gorm.ErrRecordNotFound) {
		return user, nil
	}

	return nil, err
}

func (userRepo) FindUserByPhone(ctx context.Context, phone string, status int) (*model.User, error) {
	var user *model.User

	findQ := &model.User{Phone: phone}
	if status != -1 {
		findQ.Status = status
	}

	err := app.DB().Where(findQ).First(&user).Error
	if err == nil || errors.Is(err, gorm.ErrRecordNotFound) {
		return user, nil
	}

	return nil, err
}

func (userRepo) UpdateUserLastSeen(ctx context.Context, user *model.User) error {
	now := time.Now()
	user.LastSignInAt = &now
	err := app.DB().Save(user).Error
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
	}

	return err
}
