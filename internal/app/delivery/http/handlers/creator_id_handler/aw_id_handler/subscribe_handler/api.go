package aw_subscribe_handler

import (
	"net/http"
	"tech-db-forum/internal/app/delivery/http/handlers/base_handler"
	"tech-db-forum/internal/app/delivery/http/handlers/handler_errors"
	repository_redis "tech-db-forum/internal/app/repository/pay_token/redis"
	usecase_pay_token "tech-db-forum/internal/app/usecase/pay_token"
	usecase_subscribers "tech-db-forum/internal/app/usecase/subscribers"
	"tech-db-forum/internal/pkg/utilits/postgresql"

	"github.com/sirupsen/logrus"
)

var codesByErrorsPOST = base_handler.CodeMap{
	usecase_pay_token.InvalidUserToken: {
		http.StatusBadRequest, handler_errors.InvalidUserPayToken, logrus.WarnLevel},
	repository_redis.SetError: {
		http.StatusInternalServerError, handler_errors.InternalError, logrus.ErrorLevel},
	usecase_subscribers.SubscriptionAlreadyExists: {
		http.StatusConflict, handler_errors.UserAlreadySubscribe, logrus.ErrorLevel},
	postgresql_utilits.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.BDError, logrus.ErrorLevel},
}
var codesByErrorsDELETE = base_handler.CodeMap{
	usecase_subscribers.SubscriptionsNotFound: {
		http.StatusConflict, handler_errors.SubscribesNotFound, logrus.ErrorLevel},
	postgresql_utilits.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.BDError, logrus.ErrorLevel},
}
