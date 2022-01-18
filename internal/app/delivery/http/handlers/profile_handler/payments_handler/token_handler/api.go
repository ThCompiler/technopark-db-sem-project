package payments_handler

import (
	"net/http"
	"tech-db-forum/internal/app/delivery/http/handlers/base_handler"
	"tech-db-forum/internal/app/delivery/http/handlers/handler_errors"
	repository_redis "tech-db-forum/internal/app/repository/pay_token/redis"
	repository_payments "tech-db-forum/internal/app/repository/payments"
	"tech-db-forum/internal/pkg/utilits/postgresql"

	"github.com/sirupsen/logrus"
)

var codeByErrorGET = base_handler.CodeMap{
	repository_redis.SetError: {
		http.StatusInternalServerError, handler_errors.InternalError, logrus.ErrorLevel},
	postgresql_utilits.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.InternalError, logrus.ErrorLevel},
}

var codeByErrorPOST = base_handler.CodeMap{
	repository_payments.NotEqualPaymentAmount: {
		http.StatusBadRequest, handler_errors.NotEqualPaymentAmount, logrus.ErrorLevel},
	repository_payments.CountPaymentsByTokenError: {
		http.StatusInternalServerError, handler_errors.InternalError, logrus.ErrorLevel},
	repository_redis.InvalidStorageData: {
		http.StatusInternalServerError, handler_errors.InternalError, logrus.ErrorLevel},
	postgresql_utilits.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.InternalError, logrus.ErrorLevel},
	repository_redis.NotFound: {
		http.StatusNotFound, handler_errors.PayTokenNotFound, logrus.ErrorLevel},
}
