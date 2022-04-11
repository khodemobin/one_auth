package service

import (
	"context"
	"log"
	"strconv"

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

func (r register) RegisterRequest(ctx context.Context, phone string) error {
	// TODO send verify code
	// TODO check send limit
	user, err := r.repo.UserRepo.FindUserByPhone(ctx, phone, -1)
	if err != nil {
		panic(err)
	}
	if user.ID != 0 {
		return errors.New("phone taken before")
	}

	confirmCode, code, err := encrypt.GenerateConfirmCode(phone)
	if err != nil {
		panic("internal error, can not create token")
	}
	log.Println(code)
	return r.repo.ConfirmCodeRepo.Store(phone, confirmCode)
}

func (r register) RegisterVerify(ctx context.Context, phone string, code string) (*domain.Login, error) {
	// TODO check limits
	user, err := r.repo.UserRepo.FindUserByPhone(ctx, phone, -1)
	if err != nil {
		panic(err)
	}

	if user.ID != 0 {
		return nil, errors.New("user verified!")
	}

	confirm, err := r.repo.ConfirmCodeRepo.Find(phone)
	if err != nil {
		panic(err)
	}

	if confirm == nil || !encrypt.Check(confirm.Hash, code) {
		return nil, errors.New("confirm code is invalid")
	}

	ttl, err := strconv.Atoi(r.cfg.App.JwtTTL)
	if err != nil {
		panic(err)
	}

	token, err := r.repo.TokenRepo.CreateToken(ctx, ttl, user)
	if err != nil {
		panic("internal error, can not create token")
	}

	//if err = r.repo.UserRepo.UpdateUserLastSeen(ctx, user); err != nil {
	//	panic(err)
	//}

	// TODO add event log and back and security

	return &domain.Login{
		Token:     token.Token,
		ExpiresIn: ttl,
		UserID:    user.ID,
	}, nil
}
