package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/khodemobin/pilo/auth/app"
	"github.com/khodemobin/pilo/auth/internal/model"
)

type userRepo struct{}

func NewUserRepo() UserRepository {
	return &userRepo{}
}

func (u *userRepo) Find(ctx context.Context, column string, value string) (*model.User, error) {
	var user model.User
	err := app.DB().Where(column+" = ? ", value).First(&user).Error
	err = checkError(err)

	return &user, err
}

func (u *userRepo) FindActive(ctx context.Context, column string, value string) (*model.User, error) {
	var user *model.User

	err := app.DB().Where(column+" = ? ", value).Where("is_active", 1).First(&user).Error
	err = checkError(err)

	return user, err
}

func (u *userRepo) UpdateLastSeen(ctx context.Context, user *model.User) error {
	user.LastSignInAt.Time = time.Now()
	_, err := u.Update(ctx, user)
	return err
}

func (u *userRepo) Update(ctx context.Context, user *model.User) (*model.User, error) {
	err := app.DB().Save(user).Error
	return user, err
}

func (*userRepo) Create(ctx context.Context, user *model.User) (*model.User, error) {
	err := app.DB().Create(user).Error
	return user, err
}

func (*userRepo) ExistsByID(ctx context.Context, id uint) (bool, error) {
	var count int64
	err := app.DB().Where(&model.User{
		ID: fmt.Sprintf("%d", id),
	}).Count(&count).Error

	return count > 0, err
}

func (*userRepo) ExistsByPhone(ctx context.Context, phone string) (bool, error) {
	var count int64
	err := app.DB().Where(&model.User{
		Phone: sql.NullString{
			String: phone,
		},
	}).Count(&count).Error

	return count > 0, err
}
