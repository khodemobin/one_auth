package repository

import (
	"context"
	"errors"
	"time"

	"github.com/khodemobin/pilo/auth/pkg/cache"

	"github.com/khodemobin/pilo/auth/internal/domain"
	"gorm.io/gorm"
)

type userRepo struct {
	db    *gorm.DB
	cache cache.Cache
}

func NewUserRepo(db *gorm.DB, cache cache.Cache) domain.UserRepository {
	return &userRepo{
		db:    db,
		cache: cache,
	}
}

func (u userRepo) FindUserById(ctx context.Context, id int, status int) (*domain.User, error) {
	var user *domain.User

	findQ := &domain.User{ID: uint(id)}
	if status != -1 {
		findQ.Status = status
	}

	err := u.db.Where(findQ).First(&user).Error
	if err == nil || errors.Is(err, gorm.ErrRecordNotFound) {
		return user, nil
	}

	return nil, err
}

func (u userRepo) FindUserByPhone(ctx context.Context, phone string, status int) (*domain.User, error) {
	var user *domain.User

	findQ := &domain.User{Phone: phone}
	if status != -1 {
		findQ.Status = status
	}

	err := u.db.Where(findQ).First(&user).Error
	if err == nil || errors.Is(err, gorm.ErrRecordNotFound) {
		return user, nil
	}

	return nil, err
}

func (u userRepo) UpdateUserLastSeen(ctx context.Context, user *domain.User) error {
	now := time.Now()
	user.LastSignInAt = &now
	err := u.db.Save(user).Error
	return err
}

func (u userRepo) CreateOrUpdateUser(ctx context.Context, user *domain.User) error {
	newUser := &domain.User{}

	err := u.db.Where(domain.User{Phone: user.Phone}).First(&newUser).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if err != nil {
		err = u.db.Create(&user).Error
	} else {
		err = u.db.Model(&newUser).Updates(user).Error
	}

	return err
}
