package repository_factory

import (
	"github.com/sirupsen/logrus"
	"tech-db-forum/internal/app"
	repForum "tech-db-forum/internal/app/forum/repository"
	repForumPsql "tech-db-forum/internal/app/forum/repository/postgresql"
	repPost "tech-db-forum/internal/app/post/repository"
	repPostPsql "tech-db-forum/internal/app/post/repository/postgresql"
	repService "tech-db-forum/internal/app/service/repository"
	repServicePsql "tech-db-forum/internal/app/service/repository/postgresql"
	repThread "tech-db-forum/internal/app/thread/repository"
	repThreadPsql "tech-db-forum/internal/app/thread/repository/postgresql"
	repUser "tech-db-forum/internal/app/user/repository"
	repUserPsql "tech-db-forum/internal/app/user/repository/postgresql"
)

type RepositoryFactory struct {
	expectedConnections app.ExpectedConnections
	logger              *logrus.Logger
	userRepository      repUser.Repository
	forumRepository     repForum.Repository
	postRepository      repPost.Repository
	threadRepository    repThread.Repository
	serviceRepository   repService.Repository
}

func NewRepositoryFactory(logger *logrus.Logger, expectedConnections app.ExpectedConnections) *RepositoryFactory {
	return &RepositoryFactory{
		expectedConnections: expectedConnections,
		logger:              logger,
	}
}

func (f *RepositoryFactory) GetUserRepository() repUser.Repository {
	if f.userRepository == nil {
		f.userRepository = repUserPsql.NewUserRepository(f.expectedConnections.SqlConnection)
	}
	return f.userRepository
}

func (f *RepositoryFactory) GetForumRepository() repForum.Repository {
	if f.forumRepository == nil {
		f.forumRepository = repForumPsql.NewForumRepository(f.expectedConnections.SqlConnection)
	}
	return f.forumRepository
}

func (f *RepositoryFactory) GetPostRepository() repPost.Repository {
	if f.postRepository == nil {
		f.postRepository = repPostPsql.NewPostRepository(f.expectedConnections.SqlConnection)
	}
	return f.postRepository
}

func (f *RepositoryFactory) GetServiceRepository() repService.Repository {
	if f.serviceRepository == nil {
		f.serviceRepository = repServicePsql.NewServiceRepository(f.expectedConnections.SqlConnection)
	}
	return f.serviceRepository
}

func (f *RepositoryFactory) GetThreadRepository() repThread.Repository {
	if f.threadRepository == nil {
		f.threadRepository = repThreadPsql.NewThreadRepository(f.expectedConnections.SqlConnection)
	}
	return f.threadRepository
}
