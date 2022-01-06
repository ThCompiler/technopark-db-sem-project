package upl_cover_posts_handler

import (
	"net/http"
	"tech-db-forum/internal/app/delivery/http/handlers/base_handler"
	"tech-db-forum/internal/app/delivery/http/handlers/handler_errors"
	"tech-db-forum/internal/app/repository"
	repository_os "tech-db-forum/internal/microservices/files/files/repository/files/os"
	"tech-db-forum/pkg/utils"

	"github.com/sirupsen/logrus"
)

var codeByErrorPUT = base_handler.CodeMap{
	repository.NotFound: {
		http.StatusNotFound, handler_errors.PostNotFound, logrus.ErrorLevel},
	repository.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.BDError, logrus.ErrorLevel},
	repository_os.ErrorCreate: {
		http.StatusInternalServerError, handler_errors.InternalError, logrus.ErrorLevel},
	repository_os.ErrorCopyFile: {
		http.StatusInternalServerError, handler_errors.InternalError, logrus.ErrorLevel},
	utils.ConvertErr: {
		http.StatusInternalServerError, handler_errors.InternalError, logrus.ErrorLevel},
	utils.UnknownExtOfFileName: {
		http.StatusInternalServerError, handler_errors.InternalError, logrus.ErrorLevel},
}
