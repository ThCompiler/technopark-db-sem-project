package creator_payments_handler

import (
	"net/http"
	"tech-db-forum/internal/app/delivery/http/handlers/base_handler"
	"tech-db-forum/internal/app/delivery/http/handlers/handler_errors"
	"tech-db-forum/internal/pkg/utilits/postgresql"

	"github.com/sirupsen/logrus"
)

var codeByErrorGET = base_handler.CodeMap{
	postgresql_utilits.NotFound: {
		http.StatusNoContent, handler_errors.CreatorPaymentsNotFound, logrus.WarnLevel},
	postgresql_utilits.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.InternalError, logrus.ErrorLevel},
}
