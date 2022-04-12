package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/khodemobin/pilo/auth/internal/config"
	"github.com/khodemobin/pilo/auth/internal/domain"
	"github.com/khodemobin/pilo/auth/internal/repository"
	"github.com/khodemobin/pilo/auth/pkg/helper"
	"github.com/khodemobin/pilo/auth/pkg/messenger"
)

type login struct {
	repo      *repository.Repository
	messenger messenger.Messenger
	cfg       *config.Config
}

type rabbitData struct {
	Meta   *domain.MetaData
	UserID int
	Date   time.Time
}

func NewLoginService(repo *repository.Repository, messenger messenger.Messenger, cfg *config.Config) domain.LoginService {
	return &login{
		repo:      repo,
		messenger: messenger,
		cfg:       cfg,
	}
}

func (l login) Login(ctx context.Context, phone, password string, meta *domain.MetaData) (*domain.Login, error) {
	user, err := l.repo.UserRepo.FindUserByPhone(ctx, phone, domain.USER_STATUS_ACTIVE)
	if err != nil {
		panic(fmt.Sprintf("internal error, can find user. err : %s", err.Error()))
	}

	if user.ID == 0 {
		return nil, errors.New("invalid credentials")
	}

	// if !encrypt.Check(*user.Password, password) {
	// 	return nil, errors.New("invalid credentials")
	// }

	ttl, err := strconv.Atoi(l.cfg.App.JwtTTL)
	if err != nil {
		panic(err)
	}

	token, err := l.repo.TokenRepo.CreateToken(ctx, ttl, user)
	if err != nil {
		panic(fmt.Sprintf("internal error, can not create token. err : %s", err.Error()))
	}

	// l.createMeta(int(user.ID), meta)

	return &domain.Login{
		Token:     token.Token,
		ExpiresIn: ttl,
		ID:        user.UUID,
	}, nil
}

func (login) Logout(ctx context.Context, token string, meta *domain.MetaData) error {
	return nil
}

func (l login) createMeta(userId int, meta *domain.MetaData) {
	data := &rabbitData{
		Meta:   meta,
		UserID: userId,
		Date:   time.Now(),
	}

	json, err := helper.ToJson(data)
	if err != nil {
		panic(fmt.Sprintf("internal error, can not create token. err : %s", err.Error()))
	}

	err = l.messenger.Write(json, "auth_login")
	if err != nil {
		panic(fmt.Sprintf("internal error, can not marshal rabbit data. err : %s", err.Error()))
	}
}
