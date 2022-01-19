package forum_create_handler

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"tech-db-forum/internal/app/forum/repository"
	"tech-db-forum/internal/pkg/handler/handler_errors"
	"tech-db-forum/internal/pkg/utilits/delivery"
	"tech-db-forum/internal/pkg/utilits/postgresql"
)

var codesByErrorsPOST = delivery.CodeMap{
	postgresql_utilits.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.BDError, logrus.ErrorLevel},
	repository.UserNotFound: {
		http.StatusNotFound, repository.UserNotFound, logrus.WarnLevel},
}
