package password_handler

import (
	"net/http"
	"tech-db-forum/internal/app"
	"tech-db-forum/internal/app/delivery/http/handlers/base_handler"
	"tech-db-forum/internal/app/delivery/http/handlers/handler_errors"
	"tech-db-forum/internal/app/models"
	"tech-db-forum/internal/app/repository"
	usercase_user "tech-db-forum/internal/app/usecase/user"

	log "github.com/sirupsen/logrus"
)

var codeByError = base_handler.CodeMap{
	repository.NotFound:                    {http.StatusNotFound, handler_errors.UserNotFound, log.WarnLevel},
	usercase_user.IncorrectNewPassword:     {http.StatusBadRequest, handler_errors.IncorrectNewPassword, log.InfoLevel},
	models.EmptyPassword:                   {http.StatusBadRequest, handler_errors.IncorrectNewPassword, log.InfoLevel},
	repository.DefaultErrDB:                {http.StatusInternalServerError, handler_errors.BDError, log.ErrorLevel},
	usercase_user.BadEncrypt:               {http.StatusInternalServerError, handler_errors.InternalError, log.ErrorLevel},
	app.UnknownError:                       {http.StatusInternalServerError, handler_errors.InternalError, log.ErrorLevel},
	usercase_user.OldPasswordEqualNew:      {http.StatusConflict, handler_errors.IncorrectNewPassword, log.WarnLevel},
	usercase_user.IncorrectEmailOrPassword: {http.StatusForbidden, handler_errors.IncorrectLoginOrPassword, log.InfoLevel},
}
