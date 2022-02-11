package repository_test

import (
	"context"
	"github.com/khodemobin/pilo/auth/internal/domain"
	"github.com/khodemobin/pilo/auth/internal/repository"
	"github.com/khodemobin/pilo/auth/pkg/test_mock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func Test_Repo_CreateToken(t *testing.T) {
	user, repo, db := initFakeToken(t)

	t.Run("test create token for user", func(t *testing.T) {
		token, err := repo.CreateToken(context.Background(), 5600, user)
		assert.NoError(t, err)

		foundToken := domain.Token{}
		err = db.Where(&domain.Token{ID: token.ID}).First(&foundToken).Error

		assert.NoError(t, err)
		assert.Equal(t, foundToken.Token, token.Token)

		db.Delete(&token)
	})

	db.Delete(&user)
}

func initFakeToken(t *testing.T) (*domain.User, domain.TokenRepository, *gorm.DB) {
	db, cache, _ := test_mock.NewMock(t)
	user, _ := domain.User{}.SeedUser()
	err := db.Create(user).Error
	repo := repository.NewTokenRepo(db, cache)
	assert.NoError(t, err)

	return user, repo, db
}
