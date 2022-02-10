package repository_test

import (
	"context"
	"github.com/khodemobin/pilo/auth/internal/domain"
	"github.com/khodemobin/pilo/auth/internal/repository"
	"github.com/khodemobin/pilo/auth/pkg/test_mock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestFindUserById(t *testing.T) {
	user, repo, db := initFake(t)

	t.Run("test find right user by id", func(t *testing.T) {
		u, err := repo.FindUserById(context.Background(), int(user.ID))
		assert.NoError(t, err)
		assert.Equal(t, u.ID, user.ID)
	})

	t.Run("test not found user by wrong id", func(t *testing.T) {
		u, err := repo.FindUserById(context.Background(), 123123)
		assert.Empty(t, u)
		assert.NoError(t, err)
	})

	db.Delete(&user)
}

func TestFindUserByPhone(t *testing.T) {
	user, repo, db := initFake(t)

	t.Run("test find right user by phone", func(t *testing.T) {
		u, err := repo.FindUserByPhone(context.Background(), user.Phone)
		assert.NoError(t, err)
		assert.Equal(t, u.ID, user.ID)
	})

	t.Run("test not found user by wrong phone", func(t *testing.T) {
		u, err := repo.FindUserByPhone(context.Background(), "123123")
		assert.Empty(t, u)
		assert.NoError(t, err)
	})

	db.Delete(&user)
}

func TestUpdateUserLastSeen(t *testing.T) {
	user, repo, db := initFake(t)
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

func initFake(t *testing.T) (*domain.User, domain.UserRepository, *gorm.DB) {
	db, cache, _ := test_mock.NewMock(t)
	user, _ := domain.User{}.SeedUser()
	err := db.Create(user).Error
	repo := repository.NewUserRepo(db, cache)
	assert.NoError(t, err)

	return user, repo, db
}
