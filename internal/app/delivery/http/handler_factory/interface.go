package handler_factory

import (
	useCsrf "tech-db-forum/internal/app/csrf/usecase"
	useAttaches "tech-db-forum/internal/app/usecase/attaches"
	useAwards "tech-db-forum/internal/app/usecase/awards"
	useComments "tech-db-forum/internal/app/usecase/comments"
	useCreator "tech-db-forum/internal/app/usecase/creator"
	useInfo "tech-db-forum/internal/app/usecase/info"
	useLikes "tech-db-forum/internal/app/usecase/likes"
	usePayToken "tech-db-forum/internal/app/usecase/pay_token"
	usePayments "tech-db-forum/internal/app/usecase/payments"
	usePosts "tech-db-forum/internal/app/usecase/posts"
	useStats "tech-db-forum/internal/app/usecase/statistics"
	useSubscr "tech-db-forum/internal/app/usecase/subscribers"
	useUser "tech-db-forum/internal/app/usecase/user"
)

//go:generate mockgen -destination=mocks/mock_usecase_factory.go -package=mock_usecase_factory . UsecaseFactory

type UsecaseFactory interface {
	GetUserUsecase() useUser.Usecase
	GetCreatorUsecase() useCreator.Usecase
	GetCsrfUsecase() useCsrf.Usecase
	GetAwardsUsecase() useAwards.Usecase
	GetPostsUsecase() usePosts.Usecase
	GetSubscribersUsecase() useSubscr.Usecase
	GetLikesUsecase() useLikes.Usecase
	GetAttachesUsecase() useAttaches.Usecase
	GetPaymentsUsecase() usePayments.Usecase
	GetInfoUsecase() useInfo.Usecase
	GetCommentsUsecase() useComments.Usecase
	GetStatsUsecase() useStats.Usecase
	GetPayTokenUsecase() usePayToken.Usecase
}
