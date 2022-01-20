package postgresql

import (
	"database/sql"
	"tech-db-forum/internal/app/forum"
	"time"
)

type Thread struct {
	Id      int64
	Title   string
	Author  string
	Forum   string
	Message string
	Votes   int64
	Slug    sql.NullString
	Created time.Time
}

func (t *Thread) ConvertToBaseThread() *forum.Thread {
	slug := ""
	if t.Slug.Valid {
		slug = t.Slug.String
	}

	return &forum.Thread{
		Id:      t.Id,
		Title:   t.Title,
		Author:  t.Author,
		Forum:   t.Forum,
		Message: t.Message,
		Votes:   t.Votes,
		Slug:    slug,
		Created: t.Created,
	}
}

func ConvertFromBaseThread(t *forum.Thread) *Thread  {
	slug := sql.NullString{}
	if t.Slug == "" {
		slug.Valid = false
	} else {
		slug.Valid = true
		slug.String = t.Slug
	}

	return &Thread{
		Id:      t.Id,
		Title:   t.Title,
		Author:  t.Author,
		Forum:   t.Forum,
		Message: t.Message,
		Votes:   t.Votes,
		Slug:    slug,
		Created: t.Created,
	}
}
