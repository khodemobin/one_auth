package repository

import (
	"github.com/khodemobin/pilo/auth/app"
	"github.com/khodemobin/pilo/auth/internal/model"
)

type at struct{}

func NewBlackListRepo() BlackListRepository {
	return &at{}
}

func (*at) Create(token string) error {
	return app.DB().Create(&model.BlackList{
		Token: token,
	}).Error
}

func (*at) Exists(token string) (bool, error) {
	var count int64
	err := app.DB().Where(&model.BlackList{
		Token: token,
	}).Count(&count).Error

	return count > 0, err
}
