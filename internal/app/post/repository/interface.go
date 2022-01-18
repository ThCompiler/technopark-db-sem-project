package repository

import (
	"tech-db-forum/internal/app/post"
)

type Repository interface {
	Create(posts []post.Post, threadId int64) ([]post.Post, error)
	Update(pst *post.Post) (*post.Post, error)
	SetNotEdit(id int64) (*post.Post, error)
	Get(id int64) (*post.Post, error)
	GetThreadId(slug string) (int64, error)
}
