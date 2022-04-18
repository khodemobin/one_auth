package repository_test

import (
	"context"
	"testing"

	"github.com/khodemobin/pilo/auth/internal/model"
	"github.com/khodemobin/pilo/auth/internal/repository"
	"github.com/khodemobin/pilo/auth/pkg/test_mock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func Test_Repo_CreateToken(t *testing.T) {
	user, repo, db := initFakeToken(t)

	t.Run("test create token for user", func(t *testing.T) {
		token, err := repo.CreateToken(context.Background(), user)
		assert.NoError(t, err)

		foundToken := model.RefreshToken{}
		err = db.Where(&model.RefreshToken{ID: token.ID}).First(&foundToken).Error

		assert.NoError(t, err)
		assert.Equal(t, foundToken.Token, token.Token)

		db.Delete(&token)
	})

	db.Delete(&user)
}

func initFakeToken(t *testing.T) (*model.User, repository.TokenRepository, *gorm.DB) {
	db, _, _ := test_mock.NewMock(t)
	user, _ := model.User{}.SeedUser()
	err := db.Create(user).Error
	repo := repository.NewTokenRepo()
	assert.NoError(t, err)

	return user, repo, db
}
