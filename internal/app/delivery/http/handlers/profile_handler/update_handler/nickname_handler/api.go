package nickname_handler

import (
	"net/http"
	"tech-db-forum/internal/app"
	"tech-db-forum/internal/app/delivery/http/handlers/base_handler"
	"tech-db-forum/internal/app/delivery/http/handlers/handler_errors"
	"tech-db-forum/internal/app/repository"
	usercase_user "tech-db-forum/internal/app/usecase/user"

	log "github.com/sirupsen/logrus"
)

var codeByErrorPUT = base_handler.CodeMap{
	usercase_user.InvalidOldNickname: {http.StatusUnprocessableEntity, handler_errors.InvalidOldNickname, log.WarnLevel},
	repository.NotFound:              {http.StatusNotFound, handler_errors.UserWithNicknameNotFound, log.WarnLevel},
	usercase_user.NicknameExists:     {http.StatusConflict, handler_errors.NicknameAlreadyExist, log.WarnLevel},
	repository.DefaultErrDB:          {http.StatusInternalServerError, handler_errors.BDError, log.ErrorLevel},
	app.UnknownError:                 {http.StatusInternalServerError, handler_errors.InternalError, log.ErrorLevel},
}
