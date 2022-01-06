package user_comments_handler

import (
	"net/http"
	bh "tech-db-forum/internal/app/delivery/http/handlers/base_handler"
	"tech-db-forum/internal/app/delivery/http/handlers/handler_errors"
	"tech-db-forum/internal/app/delivery/http/models"
	"tech-db-forum/internal/app/models"
	useComments "tech-db-forum/internal/app/usecase/comments"
	"tech-db-forum/internal/microservices/auth/delivery/grpc/client"
	session_middleware "tech-db-forum/internal/microservices/auth/sessions/middleware"

	"github.com/sirupsen/logrus"
)

type UserCommentsHandler struct {
	commentsUsecase useComments.Usecase
	bh.BaseHandler
}

func NewUserCommentsHandler(log *logrus.Logger,
	ucComments useComments.Usecase,
	sClient client.AuthCheckerClient) *UserCommentsHandler {
	h := &UserCommentsHandler{
		BaseHandler:     *bh.NewBaseHandler(log),
		commentsUsecase: ucComments,
	}
	sessionMiddleware := session_middleware.NewSessionMiddleware(sClient, log)

	h.AddMiddleware(sessionMiddleware.Check)
	h.AddMethod(http.MethodGet, h.GET)
	return h
}

// GET comments
// @Summary get user comments
// @tags comments
// @Description get comments for current user
// @Param page query uint64 true "start page number of posts mutually exclusive with offset"
// @Param offset query uint64 true "start number of posts mutually exclusive with page"
// @Param limit query uint64 true "posts to return"
// @Produce json
// @Success 200 {object} http_models.ResponsePostComments
// @Failure 400 {object} http_models.ErrResponse ""invalid parameters", "invalid parameters in query""
// @Failure 409 {object} http_models.ErrResponse ""comment already exist""
// @Failure 500 {object} http_models.ErrResponse ""can not do bd operation", "server error""
// @Failure 403 {object} http_models.ErrResponse ""csrf token is invalid, get new token", "this post not belongs this creators", "this user can not add comment as creator""
// @Router /user/comments [GET]
func (h *UserCommentsHandler) GET(w http.ResponseWriter, r *http.Request) {
	limit, offset, ok := h.GetPaginationFromQuery(w, r)
	if !ok {
		return
	}

	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		h.Log(r).Error("can not get user_id from context")
		h.Error(w, r, http.StatusInternalServerError, handler_errors.InternalError)
		return
	}

	res, err := h.commentsUsecase.GetUserComments(userID, &models.Pagination{Limit: limit, Offset: offset})
	if err != nil {
		h.UsecaseError(w, r, err, codesByErrorsGET)
		return
	}

	h.Log(r).Debugf("get comments for user %d", userID)
	h.Respond(w, r, http.StatusOK, http_models.ToResponseUserComments(res))
}
