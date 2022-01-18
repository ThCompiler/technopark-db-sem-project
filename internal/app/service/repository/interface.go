package repository

import "tech-db-forum/internal/app/service"

type Repository interface {
	Clear() error
	GetStat() (*service.Status, error)
}