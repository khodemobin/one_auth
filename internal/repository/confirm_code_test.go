package repository_test

import (
	"fmt"
	"testing"

	"github.com/khodemobin/pilo/auth/internal/domain"
	"github.com/khodemobin/pilo/auth/internal/repository"
	"github.com/khodemobin/pilo/auth/pkg/cache"
	"github.com/khodemobin/pilo/auth/pkg/test_mock"
	"github.com/stretchr/testify/assert"
)

func Test_Repo_CreateConfirmCode(t *testing.T) {
	user, repo, cache := initFakeConfirmCode(t)
	t.Run("test create confirm code for phone number", func(t *testing.T) {
		err := repo.CreateConfirmCode(user.Phone)
		assert.NoError(t, err)

		result, err := cache.Get(fmt.Sprintf("user_confirm_code_%s", user.Phone), nil)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func Test_Repo_FindConfirmCode(t *testing.T) {
	user, repo, _ := initFakeConfirmCode(t)

	t.Run("test find confirm code from phone number", func(t *testing.T) {
		err := repo.CreateConfirmCode(user.Phone)
		assert.NoError(t, err)

		result, err := repo.FindConfirmCode(user.Phone)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func Test_Repo_DeleteConfirmCode(t *testing.T) {
	user, repo, cache := initFakeConfirmCode(t)

	t.Run("test delete confirm code from phone number", func(t *testing.T) {
		err := repo.CreateConfirmCode(user.Phone)
		assert.NoError(t, err)

		err = repo.DeleteConfirmCode(user.Phone)
		assert.NoError(t, err)

		result, err := cache.Get(fmt.Sprintf("user_confirm_code_%s", user.Phone), nil)
		assert.NoError(t, err)
		assert.Nil(t, result)
	})
}

func initFakeConfirmCode(t *testing.T) (*domain.User, domain.ConfirmCodeRepository, cache.Cache) {
	_, cache, _ := test_mock.NewMock(t)
	user, _ := domain.User{}.SeedUser()
	repo := repository.NewConfirmCodeRepo(cache)

	return user, repo, cache
}
