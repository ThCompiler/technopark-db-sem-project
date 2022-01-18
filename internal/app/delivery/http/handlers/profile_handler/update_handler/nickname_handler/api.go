package nickname_handler

import (
	"net/http"
	"tech-db-forum/internal/app"
	"tech-db-forum/internal/app/delivery/http/handlers/base_handler"
	"tech-db-forum/internal/app/delivery/http/handlers/handler_errors"
	usercase_user "tech-db-forum/internal/app/usecase/user"
	"tech-db-forum/internal/pkg/utilits/postgresql"

	log "github.com/sirupsen/logrus"
)

var codeByErrorPUT = base_handler.CodeMap{
	usercase_user.InvalidOldNickname: {http.StatusUnprocessableEntity, handler_errors.InvalidOldNickname, log.WarnLevel},
	postgresql_utilits.NotFound:      {http.StatusNotFound, handler_errors.UserWithNicknameNotFound, log.WarnLevel},
	usercase_user.NicknameExists:     {http.StatusConflict, handler_errors.NicknameAlreadyExist, log.WarnLevel},
	postgresql_utilits.DefaultErrDB:  {http.StatusInternalServerError, handler_errors.BDError, log.ErrorLevel},
	app.UnknownError:                 {http.StatusInternalServerError, handler_errors.InternalError, log.ErrorLevel},
}
