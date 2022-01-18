package post_create_handler

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"tech-db-forum/internal/app/post/repository"
	"tech-db-forum/internal/pkg/handler/handler_errors"
	"tech-db-forum/internal/pkg/utilits/delivery"
	"tech-db-forum/internal/pkg/utilits/postgresql"
)

var codesByErrorsPOST = delivery.CodeMap{
	postgresql_utilits.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.BDError, logrus.ErrorLevel},
	repository.NotFoundPostParent: {
		http.StatusConflict, repository.NotFoundPostParent, logrus.WarnLevel},
	repository.NotFoundForumSlugOrUserOrThread: {
		http.StatusNotFound, repository.NotFoundForumSlugOrUserOrThread, logrus.WarnLevel},
}
