package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/khodemobin/pilo/auth/internal/config"
	"github.com/khodemobin/pilo/auth/internal/domain"
	"github.com/khodemobin/pilo/auth/internal/repository"
	"github.com/khodemobin/pilo/auth/internal/server/handler"
	"github.com/khodemobin/pilo/auth/pkg/encrypt"
	"github.com/khodemobin/pilo/auth/pkg/messenger"
)

type login struct {
	repo      *repository.Repository
	messenger messenger.Messenger
	cfg       *config.Config
}

type rabbitData struct {
	meta   handler.MetaData
	userID int
	date   time.Time
}

func NewLoginService(repo *repository.Repository, messenger messenger.Messenger, cfg *config.Config) domain.LoginService {
	return &login{
		repo:      repo,
		messenger: messenger,
		cfg:       cfg,
	}
}

func (l login) Login(ctx context.Context, phone, password string, meta interface{}) (*domain.Login, error) {
	user, err := l.repo.UserRepo.FindUserByPhone(ctx, phone, domain.USER_STATUS_ACTIVE)
	if err != nil {
		panic(err)
	}

	if user.ID == 0 {
		return nil, errors.New("invalid credentials")
	}

	if !encrypt.Check(*user.Password, password) {
		return nil, errors.New("invalid credentials")
	}

	ttl, err := strconv.Atoi(l.cfg.App.JwtTTL)
	if err != nil {
		panic(err)
	}

	token, err := l.repo.TokenRepo.CreateToken(ctx, ttl, user)
	if err != nil {
		panic(fmt.Sprintf("internal error, can not create token. err : %s", err.Error()))
	}

	m := meta.(handler.MetaData)
	data := &rabbitData{
		meta:   m,
		userID: int(user.ID),
		date:   time.Now(),
	}

	marshal, err := json.Marshal(data)
	if err != nil {
		panic(fmt.Sprintf("internal error, can not create token. err : %s", err.Error()))
	}

	err = l.messenger.Write(string(marshal), "auth_login")
	if err != nil {
		panic(fmt.Sprintf("internal error, can not marshal rabbit data. err : %s", err.Error()))
	}

	return &domain.Login{
		Token:     token.Token,
		ExpiresIn: ttl,
		UserID:    user.ID,
	}, nil
}

func (login) Logout(ctx context.Context, token string) error {
	return nil
}
