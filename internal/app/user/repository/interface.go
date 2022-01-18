package repository

import "tech-db-forum/internal/app/user"

type Repository interface {
	Create(us *user.User) ([]user.User, error)
	Get(nickname string) (*user.User, error)
	Update(us *user.User) (*user.User, error)
}
