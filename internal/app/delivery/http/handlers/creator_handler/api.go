package creator_handler

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"tech-db-forum/internal/app"
	"tech-db-forum/internal/app/delivery/http/handlers/base_handler"
	"tech-db-forum/internal/app/delivery/http/handlers/handler_errors"
	"tech-db-forum/internal/app/models"
	"tech-db-forum/internal/app/repository"
	repository_postgresql "tech-db-forum/internal/app/repository/creator/postgresql"
	usecase_creator "tech-db-forum/internal/app/usecase/creator"
)

var codesByErrorsGET = base_handler.CodeMap{
	repository.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.BDError, logrus.ErrorLevel},
}

var codesByErrorsPOST = base_handler.CodeMap{
	usecase_creator.CreatorExist: {
		http.StatusConflict, handler_errors.CreatorAlreadyExist, logrus.InfoLevel},
	repository.NotFound: {
		http.StatusNotFound, handler_errors.UserNotFound, logrus.WarnLevel},
	repository.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.BDError, logrus.ErrorLevel},
	app.UnknownError: {
		http.StatusInternalServerError, handler_errors.InternalError, logrus.ErrorLevel},
	models.IncorrectCreatorCategory: {
		http.StatusUnprocessableEntity, handler_errors.InvalidCategory, logrus.InfoLevel},
	repository_postgresql.IncorrectCategory: {
		http.StatusUnprocessableEntity, handler_errors.InvalidCategory, logrus.InfoLevel},
	models.IncorrectCreatorNickname: {
		http.StatusUnprocessableEntity, handler_errors.InvalidNickname, logrus.InfoLevel},
	models.IncorrectCreatorDescription: {
		http.StatusUnprocessableEntity, handler_errors.InvalidDescription, logrus.InfoLevel},
}
