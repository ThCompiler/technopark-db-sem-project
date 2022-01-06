package server

import (
	"tech-db-forum/internal/app"
)

type HandlerFactory interface {
	GetHandleUrls() *map[string]app.Handler
}
