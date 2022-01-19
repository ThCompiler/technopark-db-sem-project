package forum_threads_handler

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"tech-db-forum/internal/pkg/handler/handler_errors"
	"tech-db-forum/internal/pkg/utilits/delivery"
	"tech-db-forum/internal/pkg/utilits/postgresql"
)

var codesByErrorsGET = delivery.CodeMap{
	postgresql_utilits.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.BDError, logrus.ErrorLevel},
	postgresql_utilits.NotFound: {
		http.StatusNotFound, handler_errors.ForumNotFound, logrus.WarnLevel},
}
