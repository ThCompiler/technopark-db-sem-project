package usecase_factory

import (
	repCsrf "tech-db-forum/internal/app/csrf/repository/jwt"
	repAccess "tech-db-forum/internal/app/repository/access"
	repoAttaches "tech-db-forum/internal/app/repository/attaches"
	repoAwrds "tech-db-forum/internal/app/repository/awards"
	repoComments "tech-db-forum/internal/app/repository/comments"
	repCreator "tech-db-forum/internal/app/repository/creator"
	repoInfo "tech-db-forum/internal/app/repository/info"
	repoLikes "tech-db-forum/internal/app/repository/likes"
	repoPayToken "tech-db-forum/internal/app/repository/pay_token"
	repoPayments "tech-db-forum/internal/app/repository/payments"
	repoPosts "tech-db-forum/internal/app/repository/posts"
	repoStats "tech-db-forum/internal/app/repository/statistics"
	useSubscr "tech-db-forum/internal/app/repository/subscribers"
	repUser "tech-db-forum/internal/app/repository/user"
	push_client "tech-db-forum/internal/microservices/push/delivery/client"
)

//go:generate mockgen -destination=mocks/mock_repository_factory.go -package=mock_repository_factory . RepositoryFactory

type RepositoryFactory interface {
	GetUserRepository() repUser.Repository
	GetCreatorRepository() repCreator.Repository
	GetAwardsRepository() repoAwrds.Repository
	GetCsrfRepository() repCsrf.Repository
	GetAccessRepository() repAccess.Repository
	GetSubscribersRepository() useSubscr.Repository
	GetPostsRepository() repoPosts.Repository
	GetLikesRepository() repoLikes.Repository
	GetAttachesRepository() repoAttaches.Repository
	GetPaymentsRepository() repoPayments.Repository
	GetInfoRepository() repoInfo.Repository
	GetCommentsRepository() repoComments.Repository
	GetStatsRepository() repoStats.Repository
	GetPayTokenRepository() repoPayToken.Repository
	GetPusher() push_client.Pusher
}
