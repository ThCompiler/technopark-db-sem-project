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
					ORDER BY created, id
					LIMIT $2
					`
	getPostsFLatDESCWIthWhere = `
					SELECT id, parent, author, message, is_edited, forum, thread, created FROM posts
					WHERE thread = $1 AND id < $3
					ORDER BY created DESC, id DESC
					LIMIT $2
					`
	getPostsFLatDESCWithoutWhere = `
					SELECT id, parent, author, message, is_edited, forum, thread, created FROM posts
					WHERE thread = $1
					ORDER BY created DESC, id DESC
					LIMIT $2
					`

	getPostsThreeASCWithWhere = `
					SELECT id, parent, author, message, is_edited, forum, thread, created FROM posts
					WHERE thread = $1 AND path > (SELECT path FROM posts WHERE id = $3)
					ORDER BY path, id
					LIMIT $2
					`

	getPostsThreeASCWithoutWhere = `
					SELECT id, parent, author, message, is_edited, forum, thread, created FROM posts
					WHERE thread = $1
					ORDER BY path, id
					LIMIT $2
					`

	getPostsThreeDESCWithWhere = `
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

	getPostsParentThreeASCWIthWhere = `
					SELECT id, parent, author, message, is_edited, forum, thread, created FROM posts
					WHERE path[1] IN (
					    SELECT id FROM posts
					    WHERE thread = $1 AND parent = 0 AND path[1] > (
					        SELECT path[1] FROM posts WHERE id = $3
					        )
					    ORDER BY id 
					    LIMIT $2
					    )
					ORDER BY path, id
					`
	getPostsParentThreeASCWithoutWhere = `
					SELECT id, parent, author, message, is_edited, forum, thread, created FROM posts
					WHERE path[1] IN (
					    SELECT id FROM posts
					    WHERE thread = $1 AND parent = 0 
					    ORDER BY id 
					    LIMIT $2
					    )
					ORDER BY path, id
					`

	getPostsParentThreeDESCWIthWhere = `
					SELECT id, parent, author, message, is_edited, forum, thread, created FROM posts
					WHERE path[1] IN (
					    SELECT id FROM posts
					    WHERE thread = $1 AND parent = 0 AND path[1] < (
					        SELECT path[1] FROM posts WHERE id = $3
					        )
					    ORDER BY id DESC
					    LIMIT $2
					    )
					ORDER BY path[1] DESC, path, id
					`

	getPostsParentThreeDESCWithoutWhere = `
					SELECT id, parent, author, message, is_edited, forum, thread, created FROM posts
					WHERE path[1] IN (
					    SELECT id FROM posts
					    WHERE thread = $1 AND parent = 0
					    ORDER BY id DESC
					    LIMIT $2
					    )
					ORDER BY path[1] DESC, path, id
					`
)

func (r *ThreadRepository) getPostsFLat(id int64, pag *thread.PaginationPost) ([]thread.Post, error) {
	args := []interface{}{id, pag.Limit}
	query := getPostsFLatASC
	if pag.Desc {
		if pag.Since == app.InvalidInt {
			query = getPostsFLatDESCWithoutWhere
		} else {
			query = getPostsFLatDESCWIthWhere
			args = append(args, pag.Since)
		}
	} else {
		args = append(args, pag.Since)
	}


	rows, err := r.store.Queryx(query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, postgresql_utilits.NotFound
		}
		return nil, postgresql_utilits.NewDBError(err)
	}

	var tmp thread.Post
	res := make([]thread.Post, pag.Limit)
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
		res[i] = tmp
		i++
	}

	if i != len(res) {
		res = res[:i]
	}

	if err = rows.Err(); err != nil {
		return nil, postgresql_utilits.NewDBError(err)
	}

	return res, nil
}

func (r *ThreadRepository) getPostsThree(id int64, pag *thread.PaginationPost) ([]thread.Post, error) {
	args := []interface{}{id, pag.Limit}
	query := getPostsThreeASCWithoutWhere
	if pag.Desc {
		if pag.Since == app.InvalidInt {
			query = getPostsThreeDESCWithoutWhere
		} else {
			query = getPostsThreeDESCWithWhere
			args = append(args, pag.Since)
		}
	} else {
		if pag.Since != app.InvalidInt {
			query = getPostsThreeASCWithWhere
			args = append(args, pag.Since)
		}
	}


	rows, err := r.store.Queryx(query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, postgresql_utilits.NotFound
		}
		return nil, postgresql_utilits.NewDBError(err)
	}

	var tmp thread.Post
	res := make([]thread.Post, pag.Limit)
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
		res[i] = tmp
		i++
	}
	if i != len(res) {
		res = res[:i]
	}
	if err = rows.Err(); err != nil {
		return nil, postgresql_utilits.NewDBError(err)
	}

	return res, nil
}

func (r *ThreadRepository) getPostsParentTree(id int64, pag *thread.PaginationPost) ([]thread.Post, error) {
	args := []interface{}{id, pag.Limit}
	query := getPostsParentThreeASCWithoutWhere
	if pag.Desc {
		if pag.Since == app.InvalidInt {
			query = getPostsParentThreeDESCWithoutWhere
		} else {
			query = getPostsParentThreeDESCWIthWhere
			args = append(args, pag.Since)
		}
	} else {
		if pag.Since != app.InvalidInt {
			query = getPostsParentThreeASCWIthWhere
			args = append(args, pag.Since)
		}
	}

	rows, err := r.store.Queryx(query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, postgresql_utilits.NotFound
		}
		return nil, postgresql_utilits.NewDBError(err)
	}

	var tmp thread.Post
	var res []thread.Post
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
		res = append(res, tmp)
	}

	if err = rows.Err(); err != nil {
		return nil, postgresql_utilits.NewDBError(err)
	}

	return res, nil
}
