package comments_id_handler

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"tech-db-forum/internal/app/delivery/http/handlers/base_handler"
	"tech-db-forum/internal/app/delivery/http/handlers/handler_errors"
	"tech-db-forum/internal/pkg/utilits/postgresql"
)

var codesByErrorsPUT = base_handler.CodeMap{
	postgresql_utilits.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.BDError, logrus.ErrorLevel},
	postgresql_utilits.NotFound: {
		http.StatusNotFound, handler_errors.CommentNotFound, logrus.WarnLevel},
}

var codesByErrorsDELETE = base_handler.CodeMap{
	postgresql_utilits.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.BDError, logrus.ErrorLevel},
}
