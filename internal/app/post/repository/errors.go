package repository

import "github.com/pkg/errors"

var (
	NotFoundForumSlugOrUserOrThread = errors.New("not found forum with this slug or user or thread")
	NotFoundPostParent = errors.New("not found post parent")
)
