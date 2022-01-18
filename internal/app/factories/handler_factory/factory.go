package handler_factory

import (
	"tech-db-forum/internal/app"
	post_create_handler "tech-db-forum/internal/app/post/delivery/http/handlers/post_create"
	post_details_handler "tech-db-forum/internal/app/post/delivery/http/handlers/post_details"
	service_clear_handler "tech-db-forum/internal/app/service/delivery/http/handlers/service_clear"
	service_status_handler "tech-db-forum/internal/app/service/delivery/http/handlers/service_status"
	user_create_handler "tech-db-forum/internal/app/user/delivery/http/handlers/user_create"
	user_profile_handler "tech-db-forum/internal/app/user/delivery/http/handlers/user_profile"

	"github.com/sirupsen/logrus"
)

const (
	PROFILE = iota
	USER_CREATE
	SERVICE_CLEAR
	SERVICE_STATUS
	POST_CREATE
	POST_DETAILS
)

type HandlerFactory struct {
	repositoryFactory RepositoryFactory
	logger            *logrus.Logger
	urlHandler        *map[string]app.Handler
}

func NewFactory(logger *logrus.Logger, repositoryFactory RepositoryFactory) *HandlerFactory {
	return &HandlerFactory{
		repositoryFactory: repositoryFactory,
		logger:            logger,
	}
}

func (f *HandlerFactory) initAllHandlers() map[int]app.Handler {
	return map[int]app.Handler{
		PROFILE:        user_profile_handler.NewUserProfileHandler(f.logger, f.repositoryFactory.GetUserRepository()),
		USER_CREATE:    user_create_handler.NewUserCreateHandler(f.logger, f.repositoryFactory.GetUserRepository()),
		SERVICE_CLEAR:  service_clear_handler.NewServiceClearHandler(f.logger, f.repositoryFactory.GetServiceRepository()),
		SERVICE_STATUS: service_status_handler.NewUserProfileHandler(f.logger, f.repositoryFactory.GetServiceRepository()),
		POST_DETAILS: post_details_handler.NewPostDetailsHandler(f.logger, f.repositoryFactory.GetPostRepository(),
			f.repositoryFactory.GetThreadRepository(), f.repositoryFactory.GetUserRepository(),
			f.repositoryFactory.GetForumRepository()),
		POST_CREATE: post_create_handler.NewPostCreateHandler(f.logger, f.repositoryFactory.GetPostRepository()),
	}
}

func (f *HandlerFactory) GetHandleUrls() *map[string]app.Handler {
	if f.urlHandler != nil {
		return f.urlHandler
	}

	hs := f.initAllHandlers()
	f.urlHandler = &map[string]app.Handler{
		//=============user==============//
		"/user/<nickname>/profile": hs[PROFILE],
		"/user/<nickname>/create":  hs[USER_CREATE],

		//===========service=============//
		"/service/clear":  hs[SERVICE_CLEAR],
		"/service/status": hs[SERVICE_STATUS],

		//=============post==============//
		"/post/<id:\\d+>/details":  hs[POST_DETAILS],
		"/thread/<id:\\d+>/create": hs[POST_CREATE],
	}
	return f.urlHandler
}
