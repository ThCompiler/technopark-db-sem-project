package postgresql

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"tech-db-forum/internal/app/thread"
	"tech-db-forum/internal/app/thread/repository"
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
					UPDATE threads SET 
					    title = COALESCE(NULLIF(TRIM($1), ''), title), 
					    message = COALESCE(NULLIF(TRIM($2), ''), message) 
					WHERE id = $3 or slug = $4
					RETURNING id, title, author, forum, message, votes, slug, created
					`

	getIdBySlugQuery = "SELECT id FROM threads WHERE slug = $1"

	checkThreadQuery = "SELECT id FROM threads WHERE id = $1"

	getSlugQuery = "SELECT id, title, author, forum, message, votes, slug, created FROM threads WHERE slug = $1"
	getIdQuery   = `SELECT id, title, author, forum, message, votes, slug, created created FROM threads WHERE id = $1`

	createQuery = `    
						WITH sel AS (
						    SELECT id, title, author, forum, message, votes, slug, created
							FROM threads
							WHERE slug = $5
						), ins as (
							INSERT INTO threads (title, author, message, forum, slug, created)
								SELECT $1, $2, $3, fr.slug, $5, $6
								FROM forums  as fr
								WHERE not exists (select 1 from sel) AND fr.slug = $4
							RETURNING id, title, author, forum, message, votes, slug, created
						)
						SELECT id, title, author, forum, message, votes, slug, created, 0
						FROM ins
						UNION ALL
						SELECT id, title, author, forum, message, votes, slug, created, 1
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
	tmp := ConvertFromBaseThread(thr)

	if err := r.store.QueryRowx(createQuery,
		tmp.Title,
		tmp.Author,
		tmp.Message,
		tmp.Forum,
		tmp.Slug,
		tmp.Created,
	).
		Scan(
			&tmp.Id,
			&tmp.Title,
			&tmp.Author,
			&tmp.Forum,
			&tmp.Message,
			&tmp.Votes,
			&tmp.Slug,
			&tmp.Created,
			&isCorrect,
		); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.NotFoundForumOrAuthor
		}
		return nil, parsePQError(err.(*pq.Error))
	}
	thr = tmp.ConvertToBaseThread()

	if isCorrect == 1 {
		return thr, postgresql_utilits.Conflict
	}
	return thr, nil
}

func (r *ThreadRepository) checkThread(threadId int64) error {
	if err := r.store.QueryRowx(checkThreadQuery, threadId).Scan(&threadId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return postgresql_utilits.NotFound
		}
		return postgresql_utilits.NewDBError(err)
	}
	return nil
}

func (r *ThreadRepository) Get(id *thread.ThreadPK) (*thread.Thread, error) {
	res := Thread{}
	var err error
	if id.IsId() {
		err = r.store.QueryRowx(getIdQuery, id.GetId()).
			Scan(
				&res.Id,
				&res.Title,
				&res.Author,
				&res.Forum,
				&res.Message,
				&res.Votes,
				&res.Slug,
				&res.Created,
			)
	} else {
		err = r.store.QueryRowx(getSlugQuery, id.GetSlug()).
			Scan(
				&res.Id,
				&res.Title,
				&res.Author,
				&res.Forum,
				&res.Message,
				&res.Votes,
				&res.Slug,
				&res.Created,
			)
	}

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, postgresql_utilits.NotFound
		}
		return nil, postgresql_utilits.NewDBError(err)
	}

	return res.ConvertToBaseThread(), nil
}

func (r *ThreadRepository) GetPosts(pk *thread.ThreadPK, pag *thread.PaginationPost) ([]thread.Post, error) {
	id := int64(0)
	if !pk.IsId() {
		if err := r.store.QueryRowx(getIdBySlugQuery, pk.GetSlug()).Scan(&id); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, postgresql_utilits.NotFound
			}
			return nil, postgresql_utilits.NewDBError(err)
		}
	} else {
		id = pk.GetId()
		if err := r.checkThread(id); err != nil {
			return nil, err
		}
	}

	switch pag.Type {
	case thread.Flat:
		return r.getPostsFLat(id, pag)
	case thread.Tree:
		return r.getPostsThree(id, pag)
	case thread.ParentTree:
		return r.getPostsParentTree(id, pag)
	default:
		break
	}
	return r.getPostsFLat(id, pag)
}

func (r *ThreadRepository) Update(thr *thread.Thread) (*thread.Thread, error) {
	tmp := ConvertFromBaseThread(thr)

	if err := r.store.QueryRowx(updateQuery,
		tmp.Title,
		tmp.Message,
		tmp.Id,
		tmp.Slug,
	).Scan(
		&tmp.Id,
		&tmp.Title,
		&tmp.Author,
		&tmp.Forum,
		&tmp.Message,
		&tmp.Votes,
		&tmp.Slug,
		&tmp.Created,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, postgresql_utilits.NotFound
		}
		return nil, postgresql_utilits.NewDBError(err)
	}

	return tmp.ConvertToBaseThread(), nil
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
