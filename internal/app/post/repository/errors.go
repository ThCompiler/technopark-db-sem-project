package repository

import "github.com/pkg/errors"

var (
	NotFoundForumSlug = errors.New("not found forum with this slug")
	NotFoundPostParent = errors.New("not found post parent")
)
