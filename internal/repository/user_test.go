package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/khodemobin/pilo/auth/internal/model"
	"github.com/khodemobin/pilo/auth/internal/repository"
	"github.com/khodemobin/pilo/auth/pkg/test_mock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func Test_Repo_FindUserById(t *testing.T) {
	user, repo, db := initFakeUser(t)

	t.Run("test find right user by id", func(t *testing.T) {
		u, err := repo.FindUserById(context.Background(), int(user.ID), 1)
		assert.NoError(t, err)
		assert.Equal(t, u.ID, user.ID)
	})

	t.Run("test not found user by wrong id", func(t *testing.T) {
		u, err := repo.FindUserById(context.Background(), 123123, 0)
		assert.Empty(t, u)
		assert.NoError(t, err)
	})

	db.Delete(&user)
}

func Test_Repo_FindUserByPhone(t *testing.T) {
	user, repo, db := initFakeUser(t)

	t.Run("test find right user by phone", func(t *testing.T) {
		u, err := repo.FindUserByPhone(context.Background(), user.Phone, 1)
		assert.NoError(t, err)
		assert.Equal(t, u.ID, user.ID)
	})

	t.Run("test not found user by wrong phone", func(t *testing.T) {
		u, err := repo.FindUserByPhone(context.Background(), "123123", 0)
		assert.Empty(t, u)
		assert.NoError(t, err)
	})

	db.Delete(&user)
}

func Test_Repo_UpdateUserLastSeen(t *testing.T) {
	user, repo, db := initFakeUser(t)
	t.Run("test update user last seen after login", func(t *testing.T) {
		// last seen should be at last 2min before and 2 min later from now after update
		err := repo.UpdateUserLastSeen(context.Background(), user)
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

func initFakeUser(t *testing.T) (*model.User, repository.UserRepository, *gorm.DB) {
	db, _, _ := test_mock.NewMock(t)
	user, _ := model.User{}.SeedUser()
	err := db.Create(user).Error
	repo := repository.NewUserRepo()
	assert.NoError(t, err)

	return user, repo, db
}
