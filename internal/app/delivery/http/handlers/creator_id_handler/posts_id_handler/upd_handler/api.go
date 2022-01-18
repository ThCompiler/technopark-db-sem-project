package posts_upd_handler

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"tech-db-forum/internal/app"
	"tech-db-forum/internal/app/delivery/http/handlers/base_handler"
	"tech-db-forum/internal/app/delivery/http/handlers/handler_errors"
	"tech-db-forum/internal/app/models"
	"tech-db-forum/internal/pkg/utilits/postgresql"
)

var codesByErrorsPUT = base_handler.CodeMap{
	postgresql_utilits.NotFound: {
		http.StatusNotFound, handler_errors.PostNotFound, logrus.ErrorLevel},
	models.InvalidAwardsId: {
		http.StatusUnprocessableEntity, handler_errors.IncorrectAwardsId, logrus.InfoLevel},
	models.InvalidCreatorId: {
		http.StatusUnprocessableEntity, handler_errors.IncorrectCreatorId, logrus.WarnLevel},
	models.EmptyTitle: {
		http.StatusUnprocessableEntity, handler_errors.EmptyTitle, logrus.WarnLevel},
	postgresql_utilits.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.BDError, logrus.ErrorLevel},
	app.UnknownError: {
		http.StatusInternalServerError, handler_errors.InternalError, logrus.ErrorLevel},
}
