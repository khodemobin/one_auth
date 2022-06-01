package repository_test

import (
	"fmt"
	"testing"

	"github.com/khodemobin/pilo/auth/app"
	"github.com/khodemobin/pilo/auth/internal/model"
	"github.com/khodemobin/pilo/auth/internal/repository"
	"github.com/khodemobin/pilo/auth/pkg/test_mock"
	"github.com/stretchr/testify/assert"
)

func Test_Repo_CreateConfirmCode(t *testing.T) {
	user := initFakeConfirmCode(t)
	repo := repository.NewConfirmCodeRepo()
	t.Run("test create confirm code for phone number", func(t *testing.T) {
		err := repo.Create(*user.Phone)
		assert.NoError(t, err)

		result, err := app.Cache().Get(fmt.Sprintf("user_confirm_code_%s", *user.Phone), nil)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func Test_Repo_FindConfirmCode(t *testing.T) {
	user := initFakeConfirmCode(t)
	repo := repository.NewConfirmCodeRepo()

	t.Run("test find confirm code from phone number", func(t *testing.T) {
		err := repo.Create(*user.Phone)
		assert.NoError(t, err)

		result, err := repo.Find(*user.Phone)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func Test_Repo_DeleteConfirmCode(t *testing.T) {
	user := initFakeConfirmCode(t)
	repo := repository.NewConfirmCodeRepo()

	t.Run("test delete confirm code from phone number", func(t *testing.T) {
		err := repo.Create(*user.Phone)
		assert.NoError(t, err)

		err = repo.Delete(*user.Phone)
		assert.NoError(t, err)

		result, err := app.Cache().Get(fmt.Sprintf("user_confirm_code_%s", *user.Phone), nil)
		assert.NoError(t, err)
		assert.Nil(t, result)
	})
}

func initFakeConfirmCode(t *testing.T) *model.User {
	test_mock.NewMock(t)
	user, _ := model.User{}.SeedUser()

	return user
}
