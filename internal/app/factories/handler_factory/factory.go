package handler_factory

import (
	"tech-db-forum/internal/app"

	"github.com/sirupsen/logrus"
)

const (
	LOGIN = iota
)

type HandlerFactory struct {
	repositoryFactory    RepositoryFactory
	logger            *logrus.Logger
	urlHandler        *map[string]app.Handler
}

func NewFactory(logger *logrus.Logger, repositoryFactory RepositoryFactory) *HandlerFactory {
	return &HandlerFactory{
		repositoryFactory:    repositoryFactory,
		logger:            logger,
	}
}

func (f *HandlerFactory) initAllHandlers() map[int]app.Handler {
	return map[int]app.Handler{}
}

func (f *HandlerFactory) GetHandleUrls() *map[string]app.Handler {
	if f.urlHandler != nil {
		return f.urlHandler
	}

	hs := f.initAllHandlers()
	f.urlHandler = &map[string]app.Handler{
		"/user/<nickname>/profile": hs[LOGIN],
		"/user/<nickname>/create":  hs[LOGIN],
	}
	return f.urlHandler
}
