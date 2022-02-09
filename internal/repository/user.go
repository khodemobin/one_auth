package repository

import (
	"errors"
	"time"

	"github.com/khodemobin/pilo/auth/internal/cache"
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

func (u *userRepo) FindUserById(id int) (*domain.User, error) {
	var user *domain.User
	err := u.db.Where(&domain.User{Model: gorm.Model{ID: uint(id)}, Status: domain.USER_STATUS_ACTIVE}).First(&user).Error
	if err == nil || errors.Is(err, gorm.ErrRecordNotFound) {
		return user, nil
	}

	return nil, err
}

func (u *userRepo) FindUserByPhone(phone string) (*domain.User, error) {
	var user *domain.User
	err := u.db.Where(&domain.User{Phone: phone, Status: domain.USER_STATUS_ACTIVE}).First(&user).Error
	if err == nil || errors.Is(err, gorm.ErrRecordNotFound) {
		return user, nil
	}

	return nil, err
}

func (u *userRepo) UpdateUserLastSeen(user *domain.User) (*domain.User, error) {
	now := time.Now()
	user.LastSignInAt = &now
	err := u.db.Save(user).Error
	return user, err
}
