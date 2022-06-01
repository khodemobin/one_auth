package repository_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/khodemobin/pilo/auth/app"
	"github.com/khodemobin/pilo/auth/internal/model"
	"github.com/khodemobin/pilo/auth/internal/repository"
	"github.com/khodemobin/pilo/auth/pkg/test_mock"
	"github.com/stretchr/testify/assert"
)

func Test_Repo_Find(t *testing.T) {
	user := initFakeUser(t)
	repo := repository.NewUserRepo()
	t.Run("test find right user by id", func(t *testing.T) {
		u, err := repository.NewUserRepo().Find(context.Background(), "id", fmt.Sprint(user.ID))
		assert.NoError(t, err)
		assert.Equal(t, u.ID, user.ID)
	})

	t.Run("test find right user by uuid", func(t *testing.T) {
		u, err := repo.Find(context.Background(), "uuid", fmt.Sprint(user.UUID))
		assert.NoError(t, err)
		assert.Equal(t, u.ID, user.ID)
	})

	t.Run("test not found user by wrong id", func(t *testing.T) {
		_, err := repo.Find(context.Background(), "id", "123123")
		assert.ErrorIs(t, err, app.ErrNotFound)
	})

	app.DB().Delete(&user)
}

func Test_Repo_UpdateUserLastSeen(t *testing.T) {
	user := initFakeUser(t)
	db := app.DB()
	repo := repository.NewUserRepo()
	t.Run("test update user last seen after login", func(t *testing.T) {
		// last seen should be at last 2min before and 2 min later from now after update
		err := repo.UpdateLastSeen(context.Background(), user)
		durationAfter, _ := time.ParseDuration("-2m")
		after := time.Now().Add(durationAfter)
		durationBefore, _ := time.ParseDuration("2m")
		before := time.Now().Add(durationBefore)

		a := user.LastSignInAt.After(after)
		b := user.LastSignInAt.Before(before)

		assert.NoError(t, err)
		assert.True(t, a && b)
	})

	db.Delete(&user)
}

func initFakeUser(t *testing.T) *model.User {
	test_mock.NewMock(t)
	user, _ := model.User{}.SeedUser()
	err := app.DB().Create(user).Error

	assert.NoError(t, err)
	return user
}
