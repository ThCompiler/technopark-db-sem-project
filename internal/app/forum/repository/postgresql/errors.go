package postgresql

import (
	"github.com/lib/pq"
	"tech-db-forum/internal/app/forum/repository"
	"tech-db-forum/internal/pkg/utilits/postgresql"
)

const (
	codeIncorrectFKey = "23503"
)

func parsePQError(err *pq.Error) error {
	switch {
	case err.Code == codeIncorrectFKey:
		return repository.UserNotFound
	default:
		return postgresql_utilits.NewDBError(err)
	}
}
