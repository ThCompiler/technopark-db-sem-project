package repository

import (
	"tech-db-forum/internal/app/thread"
)

type Repository interface {
	Create(thr *thread.Thread) (*thread.Thread, error)
	Get(id *thread.ThreadPK) (*thread.Thread, error)
	GetPosts(id *thread.ThreadPK, pag *thread.PaginationPost) ([]thread.Post, error)
	Update(thr *thread.Thread) (*thread.Thread, error)
	SetVote(id *thread.ThreadPK, nickname string, value int64) (*thread.Thread, error)
}
