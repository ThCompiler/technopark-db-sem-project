package handler_factory

import (
	"tech-db-forum/internal/app"
	forum_create_handler "tech-db-forum/internal/app/forum/delivery/http/handlers/forum_create"
	forum_details_handler "tech-db-forum/internal/app/forum/delivery/http/handlers/forum_details"
	forum_threads_handler "tech-db-forum/internal/app/forum/delivery/http/handlers/forum_threads"
	forum_users_handler "tech-db-forum/internal/app/forum/delivery/http/handlers/forum_users"
	post_create_handler "tech-db-forum/internal/app/post/delivery/http/handlers/post_create"
	post_details_handler "tech-db-forum/internal/app/post/delivery/http/handlers/post_details"
	service_clear_handler "tech-db-forum/internal/app/service/delivery/http/handlers/service_clear"
	service_status_handler "tech-db-forum/internal/app/service/delivery/http/handlers/service_status"
	thread_create_handler "tech-db-forum/internal/app/thread/delivery/http/handlers/thread_create"
	thread_details_handler "tech-db-forum/internal/app/thread/delivery/http/handlers/thread_details"
	thread_posts_handler "tech-db-forum/internal/app/thread/delivery/http/handlers/thread_posts"
	thread_vote_handler "tech-db-forum/internal/app/thread/delivery/http/handlers/thread_vote"
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
	FORUM_CREATE
	FORUM_DETAILS
	FORUM_USERS
	FORUM_THREADS
	THREAD_CREATE
	THREAD_DETAILS
	THREAD_POSTS
	THREAD_VOTE
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
		SERVICE_STATUS: service_status_handler.NewServiceStatusHandler(f.logger, f.repositoryFactory.GetServiceRepository()),
		POST_DETAILS: post_details_handler.NewPostDetailsHandler(f.logger, f.repositoryFactory.GetPostRepository(),
			f.repositoryFactory.GetThreadRepository(), f.repositoryFactory.GetUserRepository(),
			f.repositoryFactory.GetForumRepository()),
		POST_CREATE:    post_create_handler.NewPostCreateHandler(f.logger, f.repositoryFactory.GetPostRepository()),
		FORUM_CREATE:   forum_create_handler.NewForumCreateHandler(f.logger, f.repositoryFactory.GetForumRepository()),
		FORUM_DETAILS:  forum_details_handler.NewForumDetailsHandler(f.logger, f.repositoryFactory.GetForumRepository()),
		FORUM_USERS:    forum_users_handler.NewForumUsersHandler(f.logger, f.repositoryFactory.GetForumRepository()),
		FORUM_THREADS:  forum_threads_handler.NewForumThreadsHandler(f.logger, f.repositoryFactory.GetForumRepository()),
		THREAD_CREATE:  thread_create_handler.NewThreadCreateHandler(f.logger, f.repositoryFactory.GetThreadRepository()),
		THREAD_DETAILS: thread_details_handler.NewThreadDetailsHandler(f.logger, f.repositoryFactory.GetThreadRepository()),
		THREAD_POSTS:   thread_posts_handler.NewThreadPostsHandler(f.logger, f.repositoryFactory.GetThreadRepository()),
		THREAD_VOTE:    thread_vote_handler.NewThreadVoteHandler(f.logger, f.repositoryFactory.GetThreadRepository()),
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
		"/thread/<slug>/create": hs[POST_CREATE],

		//=============forum=============//
		"/forum/create":         hs[FORUM_CREATE],
		"/forum/<slug>/details": hs[FORUM_DETAILS],
		"/forum/<slug>/users":   hs[FORUM_USERS],
		"/forum/<slug>/threads": hs[FORUM_THREADS],

		//============thread=============//
		"/forum/<slug>/create":   hs[THREAD_CREATE],
		"/thread/<slug>/details": hs[THREAD_DETAILS],
		"/thread/<slug>/posts":   hs[THREAD_POSTS],
		"/thread/<slug>/vote":    hs[THREAD_VOTE],
	}
	return f.urlHandler
}
