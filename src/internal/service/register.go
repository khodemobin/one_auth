package service

import (
	"context"
	"fmt"
	"time"

	"github.com/go-errors/errors"
	"github.com/khodemobin/pilo/auth/internal/model"
	"github.com/khodemobin/pilo/auth/internal/repository"
	"github.com/khodemobin/pilo/auth/pkg/encrypt"
)

type register struct {
	repo *repository.Repository
}

func NewRegisterService(repo *repository.Repository) RegisterService {
	return &register{
		repo: repo,
	}
}

func (r *register) RegisterRequest(ctx context.Context, phone string, ac *model.Activity) error {
	// TODO send verify code
	// TODO check send limit
	user, err := r.repo.UserRepo.FindUserByPhone(ctx, phone, -1)
	if err != nil {
		panic(err)
	}
	if user.ID != 0 {
		return errors.New("phone taken before")
	}

	err = r.repo.ConfirmCodeRepo.CreateConfirmCode(phone)
	if err := r.repo.ActivityRepos.CreateActivity(ac); err != nil {
		panic(err)
	}
	return err
}

func (r *register) RegisterVerify(ctx context.Context, phone string, code string, ac *model.Activity) (*Auth, error) {
	// TODO check limits
	user, err := r.repo.UserRepo.FindUserByPhone(ctx, phone, -1)
	if err != nil {
		panic(fmt.Sprintf("internal error, can not find user. err : %s", err.Error()))
	}

	if user.ID != 0 {
		return nil, errors.New("user verified!")
	}

	if err := r.checkConfirmCode(phone, code); err != nil {
		return nil, err
	}

	user = r.createUser(ctx, phone)
	r.repo.ConfirmCodeRepo.DeleteConfirmCode(phone)
	refreshToken, token := r.generateToken(ctx, user)

	if err := r.repo.ActivityRepos.CreateActivity(ac); err != nil {
		panic(err)
	}

	return &Auth{
		Token: token,
		RefreshToken: model.RefreshToken{
			Token: refreshToken.Token,
		},
		ExpiresIn: 3600, // 1 hour
		ID:        user.UUID,
	}, nil
}

func (r *register) createUser(ctx context.Context, phone string) *model.User {
	lastSeen := time.Now()
	user := &model.User{
		Phone: phone,
	}
	user.Phone = phone
	user.LastSignInAt = &lastSeen
	user.Status = model.USER_STATUS_ACTIVE

	if err := r.repo.UserRepo.CreateOrUpdateUser(ctx, user); err != nil {
		panic(fmt.Sprintf("internal error, can not find token. err : %s", err.Error()))
	}

	return user
}

func (r *register) generateToken(ctx context.Context, user *model.User) (*model.RefreshToken, string) {
	refreshToken, err := r.repo.TokenRepo.CreateToken(ctx, user)
	if err != nil {
		panic(fmt.Sprintf("internal error, can not create token. err : %s", err.Error()))
	}

	token, err := encrypt.GenerateAccessToken(user)
	if err != nil {
		panic(fmt.Sprintf("internal error, can not create token. err : %s", err.Error()))
	}

	return refreshToken, token
}

func (r *register) checkConfirmCode(phone string, code string) error {
	confirm, err := r.repo.ConfirmCodeRepo.FindConfirmCode(phone)
	if err != nil {
		panic(fmt.Sprintf("internal error, can not find confirm code. err : %s", err.Error()))
	}

	if confirm == nil || !encrypt.Check(confirm.Hash, code) {
		return errors.New("confirm code is invalid")
	}

	return nil
}