package likes_handler

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"tech-db-forum/internal/app/delivery/http/handlers/base_handler"
	"tech-db-forum/internal/app/delivery/http/handlers/handler_errors"
	"tech-db-forum/internal/app/repository"
	usecase_likes "tech-db-forum/internal/app/usecase/likes"
)

var codesByErrorsDELETE = base_handler.CodeMap{
	usecase_likes.IncorrectDelLike: {
		http.StatusConflict, handler_errors.LikesAlreadyDel, logrus.WarnLevel},
	repository.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.BDError, logrus.ErrorLevel},
}

var codesByErrorsPUT = base_handler.CodeMap{
	usecase_likes.IncorrectAddLike: {
		http.StatusConflict, handler_errors.LikesAlreadyExists, logrus.WarnLevel},
	repository.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.BDError, logrus.ErrorLevel},
}
