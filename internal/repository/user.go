package repository

import (
	"context"
	"errors"
	"github.com/khodemobin/pilo/auth/pkg/cache"
	"time"

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

func (u *userRepo) FindUserById(ctx context.Context, id int) (*domain.User, error) {
	var user *domain.User
	err := u.db.Where(&domain.User{ID: uint(id), Status: domain.USER_STATUS_ACTIVE}).First(&user).Error
	if err == nil || errors.Is(err, gorm.ErrRecordNotFound) {
		return user, nil
	}

	return nil, err
}

func (u *userRepo) FindUserByPhone(ctx context.Context, phone string) (*domain.User, error) {
	var user *domain.User
	err := u.db.Where(&domain.User{Phone: phone, Status: domain.USER_STATUS_ACTIVE}).First(&user).Error
	if err == nil || errors.Is(err, gorm.ErrRecordNotFound) {
		return user, nil
	}

	return nil, err
}

func (u *userRepo) UpdateUserLastSeen(ctx context.Context, user *domain.User) error {
	now := time.Now()
	user.LastSignInAt = &now
	err := u.db.Save(user).Error
	return err
}
