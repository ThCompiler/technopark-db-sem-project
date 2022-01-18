package likes_handler

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"tech-db-forum/internal/app/delivery/http/handlers/base_handler"
	"tech-db-forum/internal/app/delivery/http/handlers/handler_errors"
	usecase_likes "tech-db-forum/internal/app/usecase/likes"
	"tech-db-forum/internal/pkg/utilits/postgresql"
)

var codesByErrorsDELETE = base_handler.CodeMap{
	usecase_likes.IncorrectDelLike: {
		http.StatusConflict, handler_errors.LikesAlreadyDel, logrus.WarnLevel},
	postgresql_utilits.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.BDError, logrus.ErrorLevel},
}

var codesByErrorsPUT = base_handler.CodeMap{
	usecase_likes.IncorrectAddLike: {
		http.StatusConflict, handler_errors.LikesAlreadyExists, logrus.WarnLevel},
	postgresql_utilits.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.BDError, logrus.ErrorLevel},
}
