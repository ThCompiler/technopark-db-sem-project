package login_handler

import (
	"net/http"
	"tech-db-forum/internal/app/delivery/http/handlers/base_handler"
	"tech-db-forum/internal/app/delivery/http/handlers/handler_errors"
	"tech-db-forum/internal/app/models"
	"tech-db-forum/internal/app/repository"

	"github.com/sirupsen/logrus"
)

var codesByErrors = base_handler.CodeMap{
	repository.NotFound: {
		http.StatusNotFound, handler_errors.UserNotFound, logrus.WarnLevel},
	repository.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.BDError, logrus.ErrorLevel},
	models.IncorrectEmailOrPassword: {
		http.StatusUnauthorized, handler_errors.IncorrectLoginOrPassword, logrus.InfoLevel},
}
