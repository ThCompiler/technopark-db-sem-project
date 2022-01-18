package postgresql

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"tech-db-forum/internal/app/forum"
	postgresql_utilits "tech-db-forum/internal/pkg/utilits/postgresql"
)

const (
	getUsersQueryASC = `
					SELECT ur.nickname, ur.fullname, ur.about, ur.email FROM users_to_forums
					JOIN users ur on ur.nickname = users_to_forums.nickname
					WHERE forum = $1 AND ur.nickname > $2
					ORDER BY ur.nickname
					LIMIT $3
					`

	getUsersQueryDESCWithWhere = `
					SELECT ur.nickname, ur.fullname, ur.about, ur.email FROM users_to_forums
					JOIN users ur on ur.nickname = users_to_forums.nickname
					WHERE forum = $1 AND ur.nickname < $2
					ORDER BY ur.nickname DESC
					LIMIT $3
					`

	getUsersQueryDESCWithoutWhere = `
					SELECT ur.nickname, ur.fullname, ur.about, ur.email FROM users_to_forums
					JOIN users ur on ur.nickname = users_to_forums.nickname
					WHERE forum = $1
					ORDER BY ur.nickname DESC
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

	getQuery = "SELECT title, user_nickname, posts, threads FROM forums WHERE slug = $1"

	createQuery = `    
						WITH sel AS (
						    SELECT slug, title, user_nickname, posts, threads
							FROM forums
							WHERE slug = $1
						), ins as (
							INSERT INTO forums (slug, title, user_nickname)
								SELECT $1, $2, $3
								WHERE not exists (select 1 from sel)
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

func (r *ForumRepository) GetThreads(slug string, pag *forum.PaginationThread) ([]forum.Thread, error) {
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

	var res []forum.Thread
	if err := r.store.Select(&res, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, postgresql_utilits.NotFound
		}
		return nil, postgresql_utilits.NewDBError(err)
	}
	return res, nil
}

func (r *ForumRepository) GetUsers(slug string, pag *forum.PaginationUser) ([]forum.User, error) {
	query := getUsersQueryASC
	args := []interface{}{slug}
	if pag.Desc {
		if pag.Since != "" {
			query = getUsersQueryDESCWithWhere
			args = append(args, pag.Since)
		} else {
			query = getUsersQueryDESCWithoutWhere
		}
	}
	args = append(args, pag.Limit)

	var res []forum.User
	if err := r.store.Select(&res, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, postgresql_utilits.NotFound
		}
		return nil, postgresql_utilits.NewDBError(err)
	}
	return res, nil
}

func (r *ForumRepository) Get(slug string) (*forum.Forum, error) {
	res := &forum.Forum{Slug: slug}
	if err := r.store.QueryRowx(getQuery, slug).
		Scan(
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
		return nil, parsePQError(err.(*pq.Error))
	}

	if isCorrect == 1 {
		return frm, postgresql_utilits.NotFound
	}

	return frm, nil
}
