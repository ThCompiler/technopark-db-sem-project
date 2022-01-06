package csrf_handler

import (
	"net/http"
	repository_jwt "tech-db-forum/internal/app/csrf/repository/jwt"
	"tech-db-forum/internal/app/delivery/http/handlers/base_handler"
	"tech-db-forum/internal/app/delivery/http/handlers/handler_errors"

	"github.com/sirupsen/logrus"
)

var codeByErrors = base_handler.CodeMap{
	repository_jwt.ErrorSignedToken: {http.StatusInternalServerError, handler_errors.InternalError, logrus.ErrorLevel},
}
