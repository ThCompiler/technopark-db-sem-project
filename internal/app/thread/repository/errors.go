package repository

import "github.com/pkg/errors"

var (
	NotFoundForumOrAuthor = errors.New("not found forum with this slug or author")
)
