package statistics_count_posts_handler

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"tech-db-forum/internal/app/delivery/http/handlers/base_handler"
	"tech-db-forum/internal/app/delivery/http/handlers/handler_errors"
	"tech-db-forum/internal/app/repository"
	"tech-db-forum/internal/app/usecase/statistics"
)

var codeByErrorGet = base_handler.CodeMap{
	statistics.CreatorDoesNotExists: {
		http.StatusNotFound, handler_errors.CreatorNotFound, logrus.WarnLevel},
	repository.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.BDError, logrus.ErrorLevel},
}
