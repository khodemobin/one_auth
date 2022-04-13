package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-errors/errors"
	"github.com/khodemobin/pilo/auth/app"
	"github.com/khodemobin/pilo/auth/internal/model"
	"github.com/khodemobin/pilo/auth/internal/repository"
	"github.com/khodemobin/pilo/auth/pkg/helper"
)

type login struct {
	repo *repository.Repository
}

type rabbitData struct {
	Meta   *MetaData
	UserID int
	Date   time.Time
}

func NewLoginService(repo *repository.Repository) LoginService {
	return &login{
		repo: repo,
	}
}

func (l login) Login(ctx context.Context, phone, password string, meta *MetaData) (*Auth, error) {
	user, err := l.repo.UserRepo.FindUserByPhone(ctx, phone, model.USER_STATUS_ACTIVE)
	if err != nil {
		panic(fmt.Sprintf("internal error, can find user. err : %s", err.Error()))
	}

	if user.ID == 0 {
		return nil, errors.New("invalid credentials")
	}

	// if !encrypt.Check(*user.Password, password) {
	// 	return nil, errors.New("invalid credentials")
	// }

	ttl, err := strconv.Atoi(app.Config().App.JwtTTL)
	if err != nil {
		panic(err)
	}

	token, err := l.repo.TokenRepo.CreateToken(ctx, ttl, user)
	if err != nil {
		panic(fmt.Sprintf("internal error, can not create token. err : %s", err.Error()))
	}

	// l.createMeta(int(user.ID), meta)

	return &Auth{
		Token:     token.Token,
		ExpiresIn: ttl,
		ID:        user.UUID,
	}, nil
}

func (login) Logout(ctx context.Context, token string, meta *MetaData) error {
	return nil
}

func (l login) createMeta(userId int, meta *MetaData) {
	data := &rabbitData{
		Meta:   meta,
		UserID: userId,
		Date:   time.Now(),
	}

	json, err := helper.ToJson(data)
	if err != nil {
		panic(fmt.Sprintf("internal error, can not create token. err : %s", err.Error()))
	}

	err = app.Broker().Write(json, "auth_login")
	if err != nil {
		panic(fmt.Sprintf("internal error, can not marshal rabbit data. err : %s", err.Error()))
	}
}
