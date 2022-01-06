package handler_factory

import (
	useUser "tech-db-forum/internal/app/user/usecase"
)

//go:generate mockgen -destination=mocks/mock_usecase_factory.go -package=mock_usecase_factory . UsecaseFactory

type UsecaseFactory interface {
	GetUserUsecase() useUser.Usecase
}
