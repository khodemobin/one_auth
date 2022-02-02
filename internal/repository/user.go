package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/khodemobin/pilo/auth/internal/cache"
	"github.com/khodemobin/pilo/auth/internal/domain"
)

const (
	STATUS_ACTIVE    = 1
	STATUS_IN_ACTIVE = 1
	STATUS_PENDDING  = 2
)

type userRepo struct {
	db    *sqlx.DB
	cache cache.Cache
}

func NewUserRepo(db *sqlx.DB, cache cache.Cache) domain.UserRepository {
	return &userRepo{
		db:    db,
		cache: cache,
	}
}

func (u *userRepo) FindUserById(id int) (*domain.User, error) {
	return nil, nil
}

func (u *userRepo) FindUserByPhone(phone string) (*domain.User, error) {
	var user domain.User
	err := u.db.Get(&user, `
                        SELECT * FROM users
                        WHERE phone = ? AND
                        status = ?
                        LIMIT 1
                `, phone, STATUS_ACTIVE)
	if err != nil {
		return nil, err
	}

	return &user, err
}
