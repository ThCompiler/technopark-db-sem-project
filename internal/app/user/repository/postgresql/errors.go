package postgresql

import (
	"github.com/lib/pq"
	"tech-db-forum/internal/app/user/repository"
	"tech-db-forum/internal/pkg/utilits/postgresql"
)

const (
	codeDuplicateVal   = "23505"
	emailConstraint    = "users_email_key"
	nicknameConstraint = "users_pkey"
)


func parsePQError(err *pq.Error) error {
	switch {
	case err.Code == codeDuplicateVal && err.Constraint == emailConstraint:
		return repository.EmailAlreadyExist
	case err.Code == codeDuplicateVal && err.Constraint == nicknameConstraint:
		return repository.NicknameAlreadyExist
	default:
		return postgresql_utilits.NewDBError(err)
	}
}
