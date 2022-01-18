package upd_avatar_creator_handler

import (
	"net/http"
	"tech-db-forum/internal/app/delivery/http/handlers/base_handler"
	"tech-db-forum/internal/app/delivery/http/handlers/handler_errors"
	repository_os "tech-db-forum/internal/microservices/files/files/repository/files/os"
	"tech-db-forum/internal/pkg/utilits/postgresql"
	"tech-db-forum/pkg/utils"

	log "github.com/sirupsen/logrus"
)

var codeByError = base_handler.CodeMap{
	postgresql_utilits.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.BDError, log.ErrorLevel},
	postgresql_utilits.NotFound: {
		http.StatusUnprocessableEntity, handler_errors.IncorrectCreatorId, log.WarnLevel},
	repository_os.ErrorCreate: {
		http.StatusInternalServerError, handler_errors.InternalError, log.ErrorLevel},
	repository_os.ErrorCreate: {
		http.StatusInternalServerError, handler_errors.InternalError, log.ErrorLevel},
	utils.ConvertErr: {
		http.StatusInternalServerError, handler_errors.InternalError, log.ErrorLevel},
	utils.UnknownExtOfFileName: {
		http.StatusInternalServerError, handler_errors.InternalError, log.ErrorLevel},
}
