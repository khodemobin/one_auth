package service

import (
	"github.com/khodemobin/pilo/auth/internal/repository"
)

type register struct {
	repo *repository.Repository
}

//func NewRegisterService(repo *repository.Repository) RegisterService {
//	return &register{
//		repo: repo,
//	}
//}
//
//func (r *register) Request(ctx context.Context, phone string, ac *model.Activity) error {
//
//}
//
//func (r *register) Verify(ctx context.Context, phone string, code string, ac *model.Activity) (*Auth, error) {
//
//}
//
//func (r *register) createUser(ctx context.Context, phone string) *model.User {
//	lastSeen := time.Now()
//	user := &model.User{
//		Phone:        &phone,
//		LastSignInAt: &lastSeen,
//		IsActive:     true,
//	}
//
//	if _, err := r.repo.UserRepo.Create(ctx, user); err != nil {
//		panic(fmt.Sprintf("internal error, can not find token. err : %s", err.Error()))
//	}
//
//	return user
//}
//
//func (r *register) generateToken(ctx context.Context, user *model.User) (*model.RefreshToken, string) {
//	refreshToken, err := r.repo.TokenRepo.Create(ctx, user)
//	if err != nil {
//		panic(fmt.Sprintf("internal error, can not create token. err : %s", err.Error()))
//	}
//
//	token, err := encrypt.GenerateAccessToken(user)
//	if err != nil {
//		panic(fmt.Sprintf("internal error, can not create token. err : %s", err.Error()))
//	}
//
//	return refreshToken, token
//}
//
//func (r *register) checkConfirmCode(phone string, code string) error {
//	confirm, err := r.repo.ConfirmCodeRepo.Find(phone)
//	if err != nil {
//		panic(fmt.Sprintf("internal error, can not find confirm code. err : %s", err.Error()))
//	}
//
//	if confirm == nil || !encrypt.Check(confirm.Hash, code) {
//		return errors.New("confirm code is invalid")
//	}
//
//	return nil
//}
