package aw_upd_handler

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"tech-db-forum/internal/app"
	"tech-db-forum/internal/app/delivery/http/handlers/base_handler"
	"tech-db-forum/internal/app/delivery/http/handlers/handler_errors"
	"tech-db-forum/internal/app/models"
	"tech-db-forum/internal/app/repository"
	repository_postgresql "tech-db-forum/internal/app/repository/awards/postgresql"
)

var codesByErrorsPUT = base_handler.CodeMap{
	repository.NotFound: {
		http.StatusNotFound, handler_errors.AwardNotFound, logrus.ErrorLevel},
	repository.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.BDError, logrus.ErrorLevel},
	repository_postgresql.NameAlreadyExist: {
		http.StatusConflict, handler_errors.AwardsAlreadyExists, logrus.InfoLevel},
	repository_postgresql.PriceAlreadyExist: {
		http.StatusConflict, handler_errors.AwardsPriceAlreadyExists, logrus.InfoLevel},
	models.EmptyName: {
		http.StatusUnprocessableEntity, handler_errors.EmptyName, logrus.WarnLevel},
	models.IncorrectAwardsPrice: {
		http.StatusUnprocessableEntity, handler_errors.IncorrectPrice, logrus.WarnLevel},
	app.UnknownError: {
		http.StatusInternalServerError, handler_errors.InternalError, logrus.ErrorLevel},
}
