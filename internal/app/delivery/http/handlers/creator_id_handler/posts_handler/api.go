package posts_handler

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"tech-db-forum/internal/app"
	"tech-db-forum/internal/app/delivery/http/handlers/base_handler"
	"tech-db-forum/internal/app/delivery/http/handlers/handler_errors"
	"tech-db-forum/internal/app/models"
	"tech-db-forum/internal/app/repository"
)

var codesByErrorsGET = base_handler.CodeMap{
	repository.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.BDError, logrus.ErrorLevel},
}

var codesByErrorsPOST = base_handler.CodeMap{
	models.InvalidAwardsId: {
		http.StatusUnprocessableEntity, handler_errors.IncorrectAwardsId, logrus.InfoLevel},
	models.InvalidCreatorId: {
		http.StatusUnprocessableEntity, handler_errors.IncorrectCreatorId, logrus.WarnLevel},
	models.EmptyTitle: {
		http.StatusUnprocessableEntity, handler_errors.EmptyTitle, logrus.WarnLevel},
	repository.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.BDError, logrus.ErrorLevel},
	app.UnknownError: {
		http.StatusInternalServerError, handler_errors.InternalError, logrus.ErrorLevel},
}
