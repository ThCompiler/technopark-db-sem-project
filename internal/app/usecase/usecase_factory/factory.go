package usecase_factory

import (
	useUser "tech-db-forum/internal/app/user/usecase"
)

type UsecaseFactory struct {
	repositoryFactory RepositoryFactory
	userUsecase       useUser.Usecase
}

func NewUsecaseFactory(repositoryFactory RepositoryFactory) *UsecaseFactory {
	return &UsecaseFactory{
		repositoryFactory: repositoryFactory,
	}
}

func (f *UsecaseFactory) GetUserUsecase() useUser.Usecase {
	if f.userUsecase == nil {
		//f.userUsecase = useUser.NewUserUsecase(f.repositoryFactory.GetUserRepository())
	}
	return f.userUsecase
}
