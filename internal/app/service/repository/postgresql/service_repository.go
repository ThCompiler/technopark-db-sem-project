package postgresql

import (
	"github.com/jmoiron/sqlx"
	"tech-db-forum/internal/app/service"
	postgresql_utilits "tech-db-forum/internal/pkg/utilits/postgresql"
)

const (
	getStatusQuery = `SELECT 
       					(SELECT count(*) FROM forums) as cnt_forums,
						(SELECT count(*) FROM posts) as cnt_posts,
						(SELECT count(*) FROM threads) as cnt_threads,
       					(SELECT count(*) FROM users) as cnt_users
					`

	clearQuery = "TRUNCATE TABLE forums, posts, threads, users_to_forums, users, votes CASCADE;"
)

type ServiceRepository struct {
	store *sqlx.DB
}

func NewServiceRepository(store *sqlx.DB) *ServiceRepository {
	return &ServiceRepository{
		store: store,
	}
}

func (r *ServiceRepository) Clear() error {
	if _, err := r.store.Exec(clearQuery); err != nil {
		return postgresql_utilits.NewDBError(err)
	}
	return nil
}

func (r *ServiceRepository) GetStat() (*service.Status, error) {
	res := &service.Status{}
	if err := r.store.QueryRowx(getStatusQuery).
		Scan(
			&res.Forum,
			&res.Post,
			&res.Thread,
			&res.User,
			); err != nil {
		return nil, postgresql_utilits.NewDBError(err)
	}
	return res, nil
}
