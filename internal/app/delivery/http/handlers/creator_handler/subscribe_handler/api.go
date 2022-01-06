package subscribe_handler

import (
	"net/http"
	"tech-db-forum/internal/app/delivery/http/handlers/base_handler"
	"tech-db-forum/internal/app/delivery/http/handlers/handler_errors"
	"tech-db-forum/internal/app/repository"

	"github.com/sirupsen/logrus"
)

var codesByErrorsGET = base_handler.CodeMap{
	repository.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.BDError, logrus.ErrorLevel},
}
