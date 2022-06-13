package test

import (
	"github.com/khodemobin/pilo/auth/app"
	"github.com/khodemobin/pilo/auth/internal/model"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func NewMock(t *testing.T) {
	os.Setenv("APP_ENV", "test")
	app.New()
}

func initFakeUser(t *testing.T) *model.User {
	NewMock(t)
	user, _ := model.User{}.SeedUser()
	err := app.DB().Create(user).Error

	assert.NoError(t, err)
	return user
}
