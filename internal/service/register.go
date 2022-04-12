package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-errors/errors"
	"github.com/khodemobin/pilo/auth/internal/config"
	"github.com/khodemobin/pilo/auth/internal/domain"
	"github.com/khodemobin/pilo/auth/internal/repository"
	"github.com/khodemobin/pilo/auth/pkg/cache"
	"github.com/khodemobin/pilo/auth/pkg/encrypt"
	"github.com/khodemobin/pilo/auth/pkg/messenger"
)

type register struct {
	repo      *repository.Repository
	messenger messenger.Messenger
	cfg       *config.Config
	cache     cache.Cache
}

func NewRegisterService(repo *repository.Repository, messenger messenger.Messenger, c cache.Cache, cfg *config.Config) domain.RegisterService {
	return &register{
		repo:      repo,
		messenger: messenger,
		cfg:       cfg,
		cache:     c,
	}
}

func (r register) RegisterRequest(ctx context.Context, phone string, meta *domain.MetaData) error {
	// TODO send verify code
	// TODO check send limit
	user, err := r.repo.UserRepo.FindUserByPhone(ctx, phone, -1)
	if err != nil {
		panic(err)
	}
	if user.ID != 0 {
		return errors.New("phone taken before")
	}

	return r.repo.ConfirmCodeRepo.CreateConfirmCode(phone)
}

func (r register) RegisterVerify(ctx context.Context, phone string, code string, meta *domain.MetaData) (*domain.Login, error) {
	// TODO check limits
	user, err := r.repo.UserRepo.FindUserByPhone(ctx, phone, -1)
	if err != nil {
		panic(fmt.Sprintf("internal error, can not find user. err : %s", err.Error()))
	}

	if user.ID != 0 {
		return nil, errors.New("user verified!")
	}

	confirm, err := r.repo.ConfirmCodeRepo.FindConfirmCode(phone)
	if err != nil {
		panic(fmt.Sprintf("internal error, can not find confirm code. err : %s", err.Error()))
	}

	if confirm == nil || !encrypt.Check(confirm.Hash, code) {
		return nil, errors.New("confirm code is invalid")
	}

	ttl, err := strconv.Atoi(r.cfg.App.JwtTTL)
	if err != nil {
		panic(fmt.Sprintf("internal error, can not convert jwt ttl to int. err : %s", err.Error()))
	}

	user = r.createUser(ctx, phone)
	r.repo.ConfirmCodeRepo.DeleteConfirmCode(phone)
	token, err := r.repo.TokenRepo.CreateToken(ctx, ttl, user)
	if err != nil {
		panic(fmt.Sprintf("internal error, can not create token. err : %s", err.Error()))
	}

	// TODO add event log and back and security
	return &domain.Login{
		Token:     token.Token,
		ExpiresIn: ttl,
		ID:        user.UUID,
	}, nil
}

func (r register) createUser(ctx context.Context, phone string) *domain.User {
	lastSeen := time.Now()
	user := &domain.User{
		Phone: phone,
	}
	user.Phone = phone
	user.LastSignInAt = &lastSeen
	user.Status = domain.USER_STATUS_ACTIVE

	if err := r.repo.UserRepo.CreateOrUpdateUser(ctx, user); err != nil {
		panic(fmt.Sprintf("internal error, can not find token. err : %s", err.Error()))
	}

	return user
}
