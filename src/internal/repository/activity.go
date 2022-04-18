package repository

import (
	"github.com/khodemobin/pilo/auth/app"
	"github.com/khodemobin/pilo/auth/internal/model"
)

type activity struct{}

func NewActivityRepo() ActivityRepository {
	return &activity{}
}

func (*activity) CreateActivity(ac *model.Activity) error {
	return app.DB().Create(ac).Error
}
