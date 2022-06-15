package repository_test

import (
	"context"
	"github.com/khodemobin/pilo/auth/test"
	"testing"

	"github.com/khodemobin/pilo/auth/app"
	"github.com/khodemobin/pilo/auth/internal/model"
	"github.com/khodemobin/pilo/auth/internal/repository"
	"github.com/stretchr/testify/assert"
)

func Test_Repo_CreateToken(t *testing.T) {
	user := initFakeToken(t)
	repo := repository.NewTokenRepo()
	db := app.DB()
	t.Run("test create token for user", func(t *testing.T) {
		token, err := repo.Create(context.Background(), user)
		assert.NoError(t, err)

		foundToken := model.RefreshToken{}
		err = db.Where(&model.RefreshToken{ID: token.ID}).First(&foundToken).Error

		assert.NoError(t, err)
		assert.Equal(t, foundToken.Token, token.Token)

		db.Delete(&token)
	})

	db.Delete(&user)
}

func initFakeToken(t *testing.T) *model.User {
	test.NewMock(t)
	user, _ := model.User{}.SeedUser()
	err := app.DB().Create(user).Error
	assert.NoError(t, err)

	return user
}
