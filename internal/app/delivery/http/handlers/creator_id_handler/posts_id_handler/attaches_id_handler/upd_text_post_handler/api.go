package upd_text_attach_handler

import (
	"net/http"
	"tech-db-forum/internal/app"
	"tech-db-forum/internal/app/delivery/http/handlers/base_handler"
	"tech-db-forum/internal/app/delivery/http/handlers/handler_errors"
	"tech-db-forum/internal/app/models"
	repository_postgresql "tech-db-forum/internal/app/repository/attaches/postgresql"
	"tech-db-forum/internal/pkg/utilits/postgresql"

	"github.com/sirupsen/logrus"
)

var codeByErrorPUT = base_handler.CodeMap{
	postgresql_utilits.NotFound: {
		http.StatusNotFound, handler_errors.AttachNotFound, logrus.ErrorLevel},
	repository_postgresql.UnknownDataFormat: {
		http.StatusUnprocessableEntity, handler_errors.IncorrectDataType, logrus.WarnLevel},
	models.InvalidType: {
		http.StatusUnprocessableEntity, handler_errors.IncorrectDataType, logrus.WarnLevel},
	models.InvalidPostId: {
		http.StatusUnprocessableEntity, handler_errors.IncorrectPostId, logrus.WarnLevel},
	postgresql_utilits.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.BDError, logrus.ErrorLevel},
	app.UnknownError: {
		http.StatusInternalServerError, handler_errors.InternalError, logrus.ErrorLevel},
}
