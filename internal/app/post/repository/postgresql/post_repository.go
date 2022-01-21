package postgresql

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"strings"
	"tech-db-forum/internal/app"
	"tech-db-forum/internal/app/post"
	"tech-db-forum/internal/app/post/repository"
	postgresql_utilits "tech-db-forum/internal/pkg/utilits/postgresql"
	"time"
)

const (
	updateQuery = `
					UPDATE posts SET message = $1, is_edited = message != $1 WHERE id = $2 
					RETURNING parent, author, is_edited, forum, thread, created
					`

	updateNotEditQuery = `
					UPDATE posts SET is_edited = false WHERE id = $1
					RETURNING parent, author, message, is_edited, forum, thread, created
					`
	getQuery = "SELECT parent, author, message, is_edited, forum, thread, created FROM posts WHERE id = $1"

	checkParentQuery = "SELECT COUNT(id) FROM posts WHERE id in (?)"

	getThreadIdQuery  = "SELECT id FROM threads WHERE slug = $1"
	getForumSlugQuery = "SELECT forum FROM threads WHERE id = $1"

	checkThreadQuery = "SELECT id FROM threads WHERE id = $1"

	createQuery    = `INSERT INTO posts (parent, author, message, forum, thread, created) VALUES `
	createQueryEnd = ` RETURNING id, parent, author, message, is_edited, forum, thread ,created`
)

type PostRepository struct {
	store *sqlx.DB
}

func NewPostRepository(store *sqlx.DB) *PostRepository {
	return &PostRepository{
		store: store,
	}
}

func (r *PostRepository) checkParent(parent []int64, threadId int64) error {
	if len(parent) == 0 {
		return nil
	}

	query, args, err := sqlx.In(checkParentQuery, parent)
	if err != nil {
		return postgresql_utilits.NewDBError(err)
	}
	query += " AND thread = ?"
	query = r.store.Rebind(query)
	var countParent int64
	args = append(args, threadId)
	if err = r.store.QueryRowx(query, args...).Scan(&countParent); err != nil {
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

func (r *PostRepository) getForumSlug(threadId int64) (string, error) {
	res := ""
	if err := r.store.QueryRowx(getForumSlugQuery, threadId).Scan(&res); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", repository.NotFoundForumSlugOrUserOrThread
		}
		return "", postgresql_utilits.NewDBError(err)
	}
	return res, nil
}

func (r *PostRepository) checkThread(threadId int64) error {
	res := ""
	if err := r.store.QueryRowx(checkThreadQuery, threadId).Scan(&res); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return repository.NotFoundForumSlugOrUserOrThread
		}
		return postgresql_utilits.NewDBError(err)
	}
	return nil
}

func (r *PostRepository) Create(posts []post.Post, threadId int64) ([]post.Post, error) {
	if len(posts) == 0 {
		if err := r.checkThread(threadId); err != nil {
			return nil, err
		}
		return posts, nil
	}

	var argsString []string
	var args []interface{}

	forumSlug, err := r.getForumSlug(threadId)
	if err != nil {
		return nil, err
	}

	created := time.Now()
	parent := map[int64]bool{}

	for _, pst := range posts {
		argsString = append(argsString, "(?, ?, ?, ?, ?, ?)")
		pst.Created = created
		pst.Thread = threadId
		pst.Forum = forumSlug

		args = append(args, pst.Parent)
		args = append(args, pst.Author)
		args = append(args, pst.Message)
		args = append(args, pst.Forum)
		args = append(args, pst.Thread)
		args = append(args, pst.Created)

		if pst.Parent != 0 {
			parent[pst.Parent] = true
		}
	}

	arrayParent := make([]int64, len(parent))
	id := 0
	for key, _ := range parent {
		arrayParent[id] = key
		id++
	}

	if err := r.checkParent(arrayParent, threadId); err != nil {
		return nil, err
	}

	query := fmt.Sprintf("%s %s %s", createQuery,
		strings.Join(argsString, ", "), createQueryEnd)
	query = r.store.Rebind(query)

	rows, err := r.store.Queryx(query, args...)
	if err != nil {
		return nil, parsePQError(err.(*pq.Error))
	}

	var tmp post.Post
	i := 0
	for rows.Next() {
		if err = rows.Scan(
			&tmp.Id,
			&tmp.Parent,
			&tmp.Author,
			&tmp.Message,
			&tmp.Is_Edited,
			&tmp.Forum,
			&tmp.Thread,
			&tmp.Created,
		); err != nil {
			_ = rows.Close()
			return nil, postgresql_utilits.NewDBError(err)
		}
		posts[i] = tmp
		i++
	}

	if err = rows.Err(); err != nil {
		return nil, postgresql_utilits.NewDBError(err)
	}

	return posts, nil
}

func (r *PostRepository) GetThreadId(slug string) (int64, error) {
	res := int64(0)
	if err := r.store.QueryRowx(getThreadIdQuery, slug).Scan(&res); err != nil {
		return app.InvalidInt, postgresql_utilits.NewDBError(err)
	}
	return res, nil
}

func (r *PostRepository) Get(id int64) (*post.Post, error) {
	res := &post.Post{Id: id}
	if err := r.store.QueryRowx(getQuery, id).
		Scan(
			&res.Parent,
			&res.Author,
			&res.Message,
			&res.Is_Edited,
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
			&pst.Is_Edited,
			&pst.Forum,
			&pst.Thread,
			&pst.Created,
		); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, postgresql_utilits.NotFound
		}
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
			&pst.Is_Edited,
			&pst.Forum,
			&pst.Thread,
			&pst.Created,
		); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, postgresql_utilits.NotFound
		}
		return nil, postgresql_utilits.NewDBError(err)
	}
	return pst, nil
}
