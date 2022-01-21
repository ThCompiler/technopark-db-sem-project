package postgresql

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"tech-db-forum/internal/app/forum"
	"tech-db-forum/internal/app/forum/repository"
	postgresql_utilits "tech-db-forum/internal/pkg/utilits/postgresql"
)

const (
	getUsersQueryASC = `
					SELECT nickname, fullname, about, email FROM users_to_forums
					WHERE forum = $1 AND nickname > $2
					ORDER BY nickname
					LIMIT $3
					`

	getUsersQueryDESCWithWhere = `
					SELECT nickname, fullname, about, email FROM users_to_forums
					WHERE forum = $1 AND nickname < $2
					ORDER BY nickname DESC
					LIMIT $3
					`

	getUsersQueryDESCWithoutWhere = `
					SELECT nickname, fullname, about, email FROM users_to_forums
					WHERE forum = $1
					ORDER BY nickname DESC
					LIMIT $2
					`

	getThreadsQueryASCWithWhere = `
					SELECT id, title, author, forum, message, votes, slug, created FROM threads
					WHERE forum = $1 AND created >= $2
					ORDER BY created
					LIMIT $3
					`

	getThreadsQueryASCWithoutWhere = `
					SELECT id, title, author, forum, message, votes, slug, created FROM threads
					WHERE forum = $1
					ORDER BY created
					LIMIT $2
					`

	getThreadsQueryDESCWithWhere = `
					SELECT id, title, author, forum, message, votes, slug, created FROM threads
					WHERE forum = $1 AND created <= $2
					ORDER BY created DESC
					LIMIT $3
					`

	getThreadsQueryDESCWithoutWhere = `
					SELECT id, title, author, forum, message, votes, slug, created FROM threads
					WHERE forum = $1
					ORDER BY created DESC
					LIMIT $2
					`

	getQuery = "SELECT slug, title, user_nickname, posts, threads FROM forums WHERE slug = $1"

	checkSlugQuery = "SELECT slug FROM forums WHERE slug = $1"

	createQuery = `    
						WITH sel AS (
						    SELECT slug, title, user_nickname, posts, threads
							FROM forums
							WHERE slug = $1
						), ins as (
							INSERT INTO forums (slug, title, user_nickname)
								SELECT $1, $2, nickname
								FROM users
								WHERE not exists (select 1 from sel) and nickname = $3
							RETURNING slug, title, user_nickname, posts, threads
						)
						SELECT slug, title, user_nickname, posts, threads, 0
						FROM ins
						UNION ALL
						SELECT slug, title, user_nickname, posts, threads, 1
						FROM sel
					`
)

type ForumRepository struct {
	store *sqlx.DB
}

func NewForumRepository(store *sqlx.DB) *ForumRepository {
	return &ForumRepository{
		store: store,
	}
}

func (r *ForumRepository) checkSlug(slug string) error {
	if err := r.store.QueryRowx(checkSlugQuery, slug).
		Scan(
			&slug,
		); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return postgresql_utilits.NotFound
		}
		return postgresql_utilits.NewDBError(err)
	}
	return nil
}

func (r *ForumRepository) GetThreads(slug string, pag *forum.PaginationThread) ([]forum.Thread, error) {
	if err := r.checkSlug(slug); err != nil {
		return nil, err
	}

	args := []interface{}{slug}
	query := getThreadsQueryASCWithoutWhere

	if pag.Since != nil && !pag.Desc {
		query = getThreadsQueryASCWithWhere
		args = append(args, pag.Since)
	}

	if pag.Desc {
		if pag.Since != nil {
			query = getThreadsQueryDESCWithWhere
			args = append(args, pag.Since)
		} else {
			query = getThreadsQueryDESCWithoutWhere
		}
	}
	args = append(args, pag.Limit)

	rows, err := r.store.Queryx(query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, postgresql_utilits.NotFound
		}
		return nil, postgresql_utilits.NewDBError(err)
	}

	var tmp Thread
	res := make([]forum.Thread, pag.Limit)
	id := 0
	for rows.Next() {
		if err = rows.Scan(
			&tmp.Id,
			&tmp.Title,
			&tmp.Author,
			&tmp.Forum,
			&tmp.Message,
			&tmp.Votes,
			&tmp.Slug,
			&tmp.Created,
		); err != nil {
			_ = rows.Close()
			return nil, postgresql_utilits.NewDBError(err)
		}
		res = append(res, *tmp.ConvertToBaseThread())
		res[id] = *tmp.ConvertToBaseThread()
		id++
	}

	if id != len(res) {
		res = res[:id]
	}

	if err = rows.Err(); err != nil {
		return nil, postgresql_utilits.NewDBError(err)
	}

	return res, nil
}

func (r *ForumRepository) GetUsers(slug string, pag *forum.PaginationUser) ([]forum.User, error) {
	if err := r.checkSlug(slug); err != nil {
		return nil, err
	}

	query := getUsersQueryASC
	args := []interface{}{slug}
	if pag.Desc {
		if pag.Since != "" {
			query = getUsersQueryDESCWithWhere
			args = append(args, pag.Since)
		} else {
			query = getUsersQueryDESCWithoutWhere
		}
	} else {
		args = append(args, pag.Since)
	}
	args = append(args, pag.Limit)

	rows, err := r.store.Queryx(query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, postgresql_utilits.NotFound
		}
		return nil, postgresql_utilits.NewDBError(err)
	}

	var tmp forum.User
	res := make([]forum.User, pag.Limit)
	id := 0
	for rows.Next() {
		if err = rows.Scan(
			&tmp.Nickname,
			&tmp.Fullname,
			&tmp.About,
			&tmp.Email,
		); err != nil {
			_ = rows.Close()
			return nil, postgresql_utilits.NewDBError(err)
		}
		res[id] = tmp
		id++
	}

	if id != len(res) {
		res = res[:id]
	}

	if err = rows.Err(); err != nil {
		return nil, postgresql_utilits.NewDBError(err)
	}

	return res, nil
}

func (r *ForumRepository) Get(slug string) (*forum.Forum, error) {
	res := &forum.Forum{Slug: slug}
	if err := r.store.QueryRowx(getQuery, slug).
		Scan(
			&res.Slug,
			&res.Title,
			&res.User,
			&res.Posts,
			&res.Threads,
		); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, postgresql_utilits.NotFound
		}
		return nil, postgresql_utilits.NewDBError(err)
	}
	return res, nil
}

func (r *ForumRepository) Create(frm *forum.Forum) (*forum.Forum, error) {
	isCorrect := 0
	if err := r.store.QueryRowx(createQuery,
		frm.Slug,
		frm.Title,
		frm.User,
	).
		Scan(
			&frm.Slug,
			&frm.Title,
			&frm.User,
			&frm.Posts,
			&frm.Threads,
			&isCorrect,
		); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.UserNotFound
		}
		return nil, postgresql_utilits.NewDBError(err)
	}

	if isCorrect == 1 {
		return frm, postgresql_utilits.Conflict
	}

	return frm, nil
}
