package postgresql

import (
	"github.com/lib/pq"
	"tech-db-forum/internal/app/post/repository"
	"tech-db-forum/internal/pkg/utilits/postgresql"
)

const (
	codeIncorrectFKey = "42830"
)

func parsePQError(err *pq.Error) error {
	switch {
	case err.Code == codeIncorrectFKey:
		return repository.NotFoundForumSlugOrUserOrThread
	default:
		return postgresql_utilits.NewDBError(err)
	}
}
