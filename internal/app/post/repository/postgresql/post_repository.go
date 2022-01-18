package postgresql

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"strings"
	"tech-db-forum/internal/app/post"
	"tech-db-forum/internal/app/post/repository"
	postgresql_utilits "tech-db-forum/internal/pkg/utilits/postgresql"
	"time"
)

const (
	updateQuery = `
					UPDATE posts SET message = $1, is_edited = true WHERE id = $2 
					RETURNING parent, author, is_edited, forum, thread, created
					`

	updateNotEditQuery = `
					UPDATE posts SET is_edited = false WHERE id = $2 
					RETURNING parent, author, message, is_edited, forum, thread, created
					`
	getQuery = "SELECT parent, author, message, is_edited, forum, thread, created FROM posts WHERE id = $1"

	checkParentQuery = "SELECT COUNT(id) FROM posts WHERE id IN (?)"
	checkForumQuery = "SELECT slug FROM forums WHERE slug = $1"

	createQuery    = `INSERT INTO posts (parent, author, message, forum, thread, created, is_edited) VALUES `
	createQueryEnd = `RETURNING id, parent, author, message, is_edited, forum, thread ,created`
)

type PostRepository struct {
	store *sqlx.DB
}

func NewPostRepository(store *sqlx.DB) *PostRepository {
	return &PostRepository{
		store: store,
	}
}

func (r *PostRepository) checkParentAndForum(parent []int64, slug string) error {
	if err := r.store.Get(&slug, checkForumQuery, slug); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return repository.NotFoundForumSlug
		}
	}

	query, args, err := sqlx.In(checkParentQuery, parent)
	if err != nil {
		return postgresql_utilits.NewDBError(err)
	}
	query = r.store.Rebind(query)
	var countParent int64
	if err = r.store.Get(&countParent, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return repository.NotFoundPostParent
		}
		return postgresql_utilits.NewDBError(err)
	}

	if countParent != int64(len(parent)) {
		return repository.NotFoundPostParent
	}

	return nil
}

func (r *PostRepository) Create(posts []post.Post) ([]post.Post, error) {
	var argsString []string
	var args []interface{}
	created := time.Now()
	var parent []int64
	for _, pst := range posts {
		argsString = append(argsString, "(?, ?, ?, ?, ?, ?, ?)")
		pst.Created = created
		args = append(args, pst.Parent)
		args = append(args, pst.Author)
		args = append(args, pst.Message)
		args = append(args, pst.Forum)
		args = append(args, pst.Thread)
		args = append(args, pst.Created)
		args = append(args, pst.IsEdited)

		parent = append(parent, pst.Parent)
	}

	if err := r.checkParentAndForum(parent, posts[0].Forum); err != nil {
		return nil, err
	}

	query := fmt.Sprintf("%s %s %s", createQuery,
		strings.Join(argsString, ", "), createQueryEnd)
	query = r.store.Rebind(query)

	if err := r.store.Select(&posts, query, args...); err != nil {
		return nil, postgresql_utilits.NewDBError(err)
	}

	return posts, nil
}

func (r *PostRepository) Get(id int64) (*post.Post, error) {
	res := &post.Post{Id: id}
	if err := r.store.QueryRowx(getQuery, id).
		Scan(
			&res.Parent,
			&res.Author,
			&res.Message,
			&res.IsEdited,
			&res.Forum,
			&res.Thread,
			&res.Created,
			); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, postgresql_utilits.NotFound
		}
		return nil, postgresql_utilits.NewDBError(err)
	}
	return res, nil
}

func (r *PostRepository) Update(pst *post.Post) (*post.Post, error) {
	if err := r.store.QueryRowx(updateQuery, pst.Message, pst.Id).
		Scan(
			&pst.Parent,
			&pst.Author,
			&pst.IsEdited,
			&pst.Forum,
			&pst.Thread,
			&pst.Created,
			); err != nil {
		return nil, postgresql_utilits.NewDBError(err)
	}
	return pst, nil
}

func (r *PostRepository) SetNotEdit(id int64) (*post.Post, error) {
	pst := &post.Post{Id: id}
	if err := r.store.QueryRowx(updateNotEditQuery, id).
		Scan(
			&pst.Parent,
			&pst.Author,
			&pst.Message,
			&pst.IsEdited,
			&pst.Forum,
			&pst.Thread,
			&pst.Created,
			); err != nil {
		return nil, postgresql_utilits.NewDBError(err)
	}
	return pst, nil
}
