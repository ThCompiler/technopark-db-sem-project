package postgresql

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"tech-db-forum/internal/app/thread"
	postgresql_utilits "tech-db-forum/internal/pkg/utilits/postgresql"
	"time"
)

const (
	updateVotesSlugQuery = `
					INSERT INTO votes (nickname, thread_id, voice) SELECT $1, id, $3 FROM threads WHERE slug = $2
					ON CONFLICT On CONSTRAINT votes_nickname_thread_id_key 
					DO UPDATE SET voice = $3
					`

	updateVotesIdQuery = `
					INSERT INTO votes (nickname, thread_id, voice) VALUES ($1, $2, $3) 
					ON CONFLICT ON CONSTRAINT votes_nickname_thread_id_key 
					DO UPDATE SET voice = $3
					`

	updateQuery = `
					UPDATE threads SET title = $1, message = $2 WHERE id = $3 or slug = $4
					RETURNING id, title, message, forum, author, slug, votes, created
					`

	getSlugByIdQuery = "SELECT slug FROM threads WHERE ID = $1"

	getSlugQuery = "SELECT id, title, message, forum, author, slug, votes, created FROM threads WHERE slug = $1"
	getIdQuery   = `SELECT id, title, message, forum, author, slug, votes, created FROM threads WHERE id = $1`

	createQuery = `    
						WITH sel AS (
						    SELECT id, title, message, forum, author, slug, votes, created
							FROM threads
							WHERE slug = $5
						), ins as (
							INSERT INTO threads (title, author, message, forum, slug)
								SELECT $1, $2, $3, $4, $5
								WHERE not exists (select 1 from sel)
							RETURNING id, title, message, forum, author, slug, votes, created
						)
						SELECT id, title, message, forum, author, slug, votes, created, 0
						FROM ins
						UNION ALL
						SELECT id, title, message, forum, author, slug, votes, created, 1
						FROM sel
					`
)

type ThreadRepository struct {
	store *sqlx.DB
}

func NewThreadRepository(store *sqlx.DB) *ThreadRepository {
	return &ThreadRepository{
		store: store,
	}
}

func (r *ThreadRepository) Create(thr *thread.Thread) (*thread.Thread, error) {
	isCorrect := 0
	if thr.Created.IsZero() {
		thr.Created = time.Now()
	}

	if err := r.store.QueryRowx(createQuery,
		thr.Title,
		thr.Author,
		thr.Message,
		thr.Forum,
		thr.Slug,
	).
		Scan(
			&thr.Id,
			&thr.Title,
			&thr.Message,
			&thr.Forum,
			&thr.Author,
			&thr.Slug,
			&thr.Votes,
			&thr.Created,
			&isCorrect,
		); err != nil {
		return nil, parsePQError(err.(*pq.Error))
	}

	if isCorrect == 1 {
		return thr, postgresql_utilits.NotFound
	}
	return thr, nil
}

func (r *ThreadRepository) Get(id *thread.ThreadPK) (*thread.Thread, error) {
	res := &thread.Thread{}
	var err error
	if id.IsId() {
		err = r.store.Get(res, getIdQuery, id.GetId())
	} else {
		err = r.store.Get(res, getSlugQuery, id.GetSlug())
	}

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, postgresql_utilits.NotFound
		}
		return nil, postgresql_utilits.NewDBError(err)
	}

	return res, nil
}

func (r *ThreadRepository) GetPosts(id *thread.ThreadPK, pag *thread.PaginationPost) ([]thread.Post, error) {
	slug := ""
	if id.IsId() {
		if err := r.store.Get(&slug, getSlugByIdQuery, id.GetId()); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, postgresql_utilits.NotFound
			}
			return nil, postgresql_utilits.NewDBError(err)
		}
	} else {
		slug = id.GetSlug()
	}

	switch pag.Type {
	case thread.Flat:
		return r.getPostsFLat(slug, pag)
	case thread.Tree:
		return r.getPostsThree(slug, pag)
	case thread.ParentTree:
		return r.getPostsParentTree(slug, pag)
	default:
		break
	}
	return r.getPostsFLat(slug, pag)
}

func (r *ThreadRepository) Update(thr *thread.Thread) (*thread.Thread, error) {
	if err := r.store.Get(thr, updateQuery,
		thr.Title,
		thr.Message,
		thr.Id,
		thr.Slug,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, postgresql_utilits.NotFound
		}
		return nil, postgresql_utilits.NewDBError(err)
	}
	return thr, nil
}

func (r *ThreadRepository) SetVote(id *thread.ThreadPK, nickname string, value int64) (*thread.Thread, error) {
	var err error
	if id.IsId() {
		_, err = r.store.Exec(updateVotesIdQuery, nickname, id.GetId(), value)
	} else {
		_, err = r.store.Exec(updateVotesSlugQuery, nickname, id.GetSlug(), value)
	}

	if err != nil {
		return nil, parsePQError(err.(*pq.Error))
	}

	return r.Get(id)
}

