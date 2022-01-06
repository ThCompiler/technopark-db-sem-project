package repository

import (
	"errors"
	"tech-db-forum/internal/app"
)

const (
	NoAwards = -1
)

var (
	DefaultErrDB = errors.New("something wrong DB")
	NotFound     = errors.New("user not found")
)

func NewDBError(externalErr error) *app.GeneralError {
	return &app.GeneralError{
		Err:         DefaultErrDB,
		ExternalErr: externalErr,
	}

}
