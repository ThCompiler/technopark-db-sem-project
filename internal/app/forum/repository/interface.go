package repository

import (
	"tech-db-forum/internal/app/forum"
)

type Repository interface {
	Create(frm *forum.Forum) (*forum.Forum, error)
	GetThreads(slug string, pag *forum.PaginationThread) ([]forum.Thread, error)
	GetUsers(slug string, pag *forum.PaginationUser) ([]forum.User, error)
	Get(slug string) (*forum.Forum, error)
}
