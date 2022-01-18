package handler_factory

import (
	repForum "tech-db-forum/internal/app/forum/repository"
	repPost "tech-db-forum/internal/app/post/repository"
	repService "tech-db-forum/internal/app/service/repository"
	repThread "tech-db-forum/internal/app/thread/repository"
	repUser "tech-db-forum/internal/app/user/repository"
)

//go:generate mockgen -destination=mocks/mock_repository_factory.go -package=mock_repository_factory . RepositoryFactory

type RepositoryFactory interface {
	GetUserRepository() repUser.Repository
	GetForumRepository() repForum.Repository
	GetPostRepository() repPost.Repository
	GetServiceRepository() repService.Repository
	GetThreadRepository() repThread.Repository
}
