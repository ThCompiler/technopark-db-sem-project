package postgresql

import (
	"database/sql"
	"github.com/pkg/errors"
	"tech-db-forum/internal/app"
	"tech-db-forum/internal/app/thread"
	"tech-db-forum/internal/pkg/utilits/postgresql"
)

const (
	getPostsFLatASC = `
					SELECT id, parent, author, message, is_edited, forum, thread, created FROM posts
					WHERE thread = $1 AND id > $3
					ORDER BY created
					LIMIT $2
					`
	getPostsFLatDESCWIthWhere = `
					SELECT id, parent, author, message, is_edited, forum, thread, created FROM posts
					WHERE thread = $1 AND id < $3
					ORDER BY created DESC
					LIMIT $2
					`
	getPostsFLatDESCWithoutWhere = `
					SELECT id, parent, author, message, is_edited, forum, thread, created FROM posts
					WHERE thread = $1
					ORDER BY created DESC
					LIMIT $2
					`

	getPostsThreeASC = `
					SELECT id, parent, author, message, is_edited, forum, thread, created FROM posts
					WHERE thread = $1 AND path > (SELECT path FROM posts WHERE id = $3)
					ORDER BY path
					LIMIT $2
					`

	getPostsThreeDESCWIthWhere = `
					SELECT id, parent, author, message, is_edited, forum, thread, created FROM posts
					WHERE thread = $1 AND path < (SELECT path FROM posts WHERE id = $3)
					ORDER BY path DESC
					LIMIT $2
					`

	getPostsThreeDESCWithoutWhere = `
					SELECT id, parent, author, message, is_edited, forum, thread, created FROM posts
					WHERE thread = $1
					ORDER BY path DESC
					LIMIT $2
					`

	getPostsParentThreeASC = `
					SELECT id, parent, author, message, is_edited, forum, thread, created FROM posts
					WHERE path[1] IN (
					    SELECT id FROM posts
					    WHERE thread = $1 AND parent = 0 AND path[1] > (
					        SELECT path FROM posts WHERE id = $3
					        )
					    ORDER BY id 
					    LIMIT $2
					    )
					ORDER BY path
					`

	getPostsParentThreeDESCWIthWhere = `
					SELECT id, parent, author, message, is_edited, forum, thread, created FROM posts
					WHERE path[1] IN (
					    SELECT id FROM posts
					    WHERE thread = $1 AND parent = 0 AND path[1] < (
					        SELECT path FROM posts WHERE id = $3
					        )
					    ORDER BY id DESC
					    LIMIT $2
					    )
					ORDER BY path[1] DESC, path
					`

	getPostsParentThreeDESCWithoutWhere = `
					SELECT id, parent, author, message, is_edited, forum, thread, created FROM posts
					WHERE path[1] IN (
					    SELECT id FROM posts
					    WHERE thread = $1 AND parent = 0
					    ORDER BY id DESC
					    LIMIT $2
					    )
					ORDER BY path[1] DESC, path
					`
)

func (r *ThreadRepository) getPostsFLat(slug string, pag *thread.PaginationPost) ([]thread.Post, error) {
	args := []interface{}{slug, pag.Limit}
	query := getPostsFLatASC
	if pag.Desc {
		if pag.Since == app.InvalidInt {
			query = getPostsFLatDESCWithoutWhere
		} else {
			query = getPostsFLatDESCWIthWhere
			args = append(args, pag.Since)
		}
	}

	var res []thread.Post
	if err := r.store.Select(&res, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, postgresql_utilits.NotFound
		}
		return nil, postgresql_utilits.NewDBError(err)
	}

	return res, nil
}

func (r *ThreadRepository) getPostsThree(slug string, pag *thread.PaginationPost) ([]thread.Post, error) {
	args := []interface{}{slug, pag.Limit}
	query := getPostsThreeASC
	if pag.Desc {
		if pag.Since == app.InvalidInt {
			query = getPostsThreeDESCWIthWhere
		} else {
			query = getPostsThreeDESCWithoutWhere
			args = append(args, pag.Since)
		}
	}

	var res []thread.Post
	if err := r.store.Select(&res, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, postgresql_utilits.NotFound
		}
		return nil, postgresql_utilits.NewDBError(err)
	}

	return res, nil
}

func (r *ThreadRepository) getPostsParentTree(slug string, pag *thread.PaginationPost) ([]thread.Post, error) {
	args := []interface{}{slug, pag.Limit}
	query := getPostsParentThreeASC
	if pag.Desc {
		if pag.Since == app.InvalidInt {
			query = getPostsParentThreeDESCWIthWhere
		} else {
			query = getPostsParentThreeDESCWithoutWhere
			args = append(args, pag.Since)
		}
	}

	var res []thread.Post
	if err := r.store.Select(&res, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, postgresql_utilits.NotFound
		}
		return nil, postgresql_utilits.NewDBError(err)
	}

	return res, nil
}
