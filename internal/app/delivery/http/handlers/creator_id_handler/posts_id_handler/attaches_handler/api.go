package attaches_handler

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"tech-db-forum/internal/app/delivery/http/handlers/base_handler"
	"tech-db-forum/internal/app/delivery/http/handlers/handler_errors"
	"tech-db-forum/internal/app/models"
	repository_postgresql "tech-db-forum/internal/app/repository/attaches/postgresql"
	"tech-db-forum/internal/pkg/utilits/postgresql"
)

var codesByErrorsPOST = base_handler.CodeMap{
	postgresql_utilits.NotFound: {
		http.StatusNotFound, handler_errors.AttachNotFound, logrus.ErrorLevel},
	postgresql_utilits.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.BDError, logrus.ErrorLevel},
	models.IncorrectType: {
		http.StatusUnprocessableEntity, handler_errors.IncorrectType, logrus.WarnLevel},
	models.IncorrectAttachId: {
		http.StatusUnprocessableEntity, handler_errors.IncorrectIdAttach, logrus.WarnLevel},
	models.IncorrectLevel: {
		http.StatusInternalServerError, handler_errors.InternalError, logrus.ErrorLevel},
	repository_postgresql.UnknownDataFormat: {
		http.StatusInternalServerError, handler_errors.IncorrectType, logrus.ErrorLevel},
}
