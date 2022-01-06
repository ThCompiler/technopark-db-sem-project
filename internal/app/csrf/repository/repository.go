package repository_token

import "tech-db-forum/internal/app/csrf/csrf_models"

type Repository interface {
	Check(sources csrf_models.TokenSources, tokenString csrf_models.Token) error
	Create(sources csrf_models.TokenSources) (csrf_models.Token, error)
}
