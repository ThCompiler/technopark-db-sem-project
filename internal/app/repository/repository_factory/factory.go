package repository_factory

import (
	"tech-db-forum/internal/app"
	repCsrf "tech-db-forum/internal/app/csrf/repository/jwt"
	repositoryAccess "tech-db-forum/internal/app/repository/access"
	repoAttaches "tech-db-forum/internal/app/repository/attaches"
	repoAttachesPsql "tech-db-forum/internal/app/repository/attaches/postgresql"
	repoAwrds "tech-db-forum/internal/app/repository/awards"
	repAwardsPsql "tech-db-forum/internal/app/repository/awards/postgresql"
	repoComments "tech-db-forum/internal/app/repository/comments"
	repCommentsPsql "tech-db-forum/internal/app/repository/comments/postgresql"
	repCreator "tech-db-forum/internal/app/repository/creator"
	repCreatorPsql "tech-db-forum/internal/app/repository/creator/postgresql"
	repoInfo "tech-db-forum/internal/app/repository/info"
	repInfoPsql "tech-db-forum/internal/app/repository/info/postgresql"
	repoLikes "tech-db-forum/internal/app/repository/likes"
	repLikesPsql "tech-db-forum/internal/app/repository/likes/postgresql"
	repoPayToken "tech-db-forum/internal/app/repository/pay_token"
	repoPayTokenRedis "tech-db-forum/internal/app/repository/pay_token/redis"
	repoPayments "tech-db-forum/internal/app/repository/payments"
	repoPaymentsPsql "tech-db-forum/internal/app/repository/payments/postgresql"
	repoPosts "tech-db-forum/internal/app/repository/posts"
	repPostsPsql "tech-db-forum/internal/app/repository/posts/postgresql"
	repStats "tech-db-forum/internal/app/repository/statistics"
	repStatsPsql "tech-db-forum/internal/app/repository/statistics/postgresql"
	repoSubscribers "tech-db-forum/internal/app/repository/subscribers"
	repUser "tech-db-forum/internal/app/repository/user"
	repUserPsql "tech-db-forum/internal/app/repository/user/postgresql"
	push_client "tech-db-forum/internal/microservices/push/delivery/client"

	"github.com/sirupsen/logrus"
)

type RepositoryFactory struct {
	expectedConnections   app.ExpectedConnections
	paymentsConfig        app.Payments
	logger                *logrus.Logger
	userRepository        repUser.Repository
	creatorRepository     repCreator.Repository
	awardsRepository      repoAwrds.Repository
	postsRepository       repoPosts.Repository
	likesRepository       repoLikes.Repository
	AttachRepository      repoAttaches.Repository
	csrfRepository        repCsrf.Repository
	accessRepository      repositoryAccess.Repository
	subscribersRepository repoSubscribers.Repository
	paymentsRepository    repoPayments.Repository
	infoRepository        repoInfo.Repository
	statsRepository       repStats.Repository
	payTokenRepository    repoPayToken.Repository
	commentsRepository    repoComments.Repository
	pusher                push_client.Pusher
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

func (f *RepositoryFactory) GetCreatorRepository() repCreator.Repository {
	if f.creatorRepository == nil {
		f.creatorRepository = repCreatorPsql.NewCreatorRepository(f.expectedConnections.SqlConnection)
	}
	return f.creatorRepository
}

func (f *RepositoryFactory) GetCsrfRepository() repCsrf.Repository {
	if f.csrfRepository == nil {
		f.csrfRepository = repCsrf.NewJwtRepository()
	}
	return f.csrfRepository
}

func (f *RepositoryFactory) GetAwardsRepository() repoAwrds.Repository {
	if f.awardsRepository == nil {
		f.awardsRepository = repAwardsPsql.NewAwardsRepository(f.expectedConnections.SqlConnection)
	}
	return f.awardsRepository
}
func (f *RepositoryFactory) GetAccessRepository() repositoryAccess.Repository {
	if f.accessRepository == nil {
		f.accessRepository = repositoryAccess.NewRedisRepository(f.expectedConnections.AccessRedisPool, f.logger)
	}
	return f.accessRepository
}
func (f *RepositoryFactory) GetSubscribersRepository() repoSubscribers.Repository {
	if f.subscribersRepository == nil {
		f.subscribersRepository = repoSubscribers.NewSubscribersRepository(f.expectedConnections.SqlConnection)
	}
	return f.subscribersRepository
}

func (f *RepositoryFactory) GetPostsRepository() repoPosts.Repository {
	if f.postsRepository == nil {
		f.postsRepository = repPostsPsql.NewPostsRepository(f.expectedConnections.SqlConnection)
	}
	return f.postsRepository
}

func (f *RepositoryFactory) GetLikesRepository() repoLikes.Repository {
	if f.likesRepository == nil {
		f.likesRepository = repLikesPsql.NewLikesRepository(f.expectedConnections.SqlConnection)
	}
	return f.likesRepository
}

func (f *RepositoryFactory) GetAttachesRepository() repoAttaches.Repository {
	if f.AttachRepository == nil {
		f.AttachRepository = repoAttachesPsql.NewAttachesRepository(f.expectedConnections.SqlConnection)
	}
	return f.AttachRepository
}

func (f *RepositoryFactory) GetPaymentsRepository() repoPayments.Repository {
	if f.paymentsRepository == nil {
		f.paymentsRepository = repoPaymentsPsql.NewPaymentsRepository(f.expectedConnections.SqlConnection)
	}
	return f.paymentsRepository
}

func (f *RepositoryFactory) GetInfoRepository() repoInfo.Repository {
	if f.infoRepository == nil {
		f.infoRepository = repInfoPsql.NewInfoRepository(f.expectedConnections.SqlConnection)
	}
	return f.infoRepository
}
func (f *RepositoryFactory) GetStatsRepository() repStats.Repository {
	if f.statsRepository == nil {
		f.statsRepository = repStatsPsql.NewStatisticsRepository(f.expectedConnections.SqlConnection)
	}
	return f.statsRepository
}

func (f *RepositoryFactory) GetCommentsRepository() repoComments.Repository {
	if f.commentsRepository == nil {
		f.commentsRepository = repCommentsPsql.NewCommentsRepository(f.expectedConnections.SqlConnection)
	}
	return f.commentsRepository
}

func (f *RepositoryFactory) GetPusher() push_client.Pusher {
	if f.pusher == nil {
		f.pusher = push_client.NewPushSender(f.expectedConnections.RabbitSession)
	}
	return f.pusher
}
func (f *RepositoryFactory) GetPayTokenRepository() repoPayToken.Repository {
	if f.payTokenRepository == nil {
		f.payTokenRepository = repoPayTokenRedis.NewPayTokenRepository(f.expectedConnections.AccessRedisPool)
	}
	return f.payTokenRepository
}
