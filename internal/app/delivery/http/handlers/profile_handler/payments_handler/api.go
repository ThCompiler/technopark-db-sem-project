package payments_handler

import (
	"net/http"
	"tech-db-forum/internal/app/delivery/http/handlers/base_handler"
	"tech-db-forum/internal/app/delivery/http/handlers/handler_errors"
	"tech-db-forum/internal/app/repository"

	"github.com/sirupsen/logrus"
)

var codeByErrorGET = base_handler.CodeMap{
	repository.NotFound: {
		http.StatusNoContent, handler_errors.PaymentsNotFound, logrus.WarnLevel},
	repository.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.InternalError, logrus.ErrorLevel},
}
