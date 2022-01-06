package postgresql

import (
	"github.com/jmoiron/sqlx"
	"tech-db-forum/internal/app/user"
)

const (
	CreateQuery = `INSERT`
)


type UserRepository struct {
	store *sqlx.DB
}

func NewUserRepository(store *sqlx.DB) *UserRepository {
	return &UserRepository{
		store: store,
	}
}

func (r *UserRepository) Create(us *user.User) (*user.User, error) {
	return nil, nil
}

func (r *UserRepository) Get(nickname string) error {
	return nil
}

func (r *UserRepository) Update(us *user.User) (*user.User, error) {
	return nil, nil
}
