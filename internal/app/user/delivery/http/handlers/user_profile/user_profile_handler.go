package user_profile_handler

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"tech-db-forum/internal/app/user"
	"tech-db-forum/internal/app/user/delivery/http"
	"tech-db-forum/internal/app/user/repository"
	bh "tech-db-forum/internal/pkg/handler"
	"tech-db-forum/internal/pkg/handler/handler_errors"
)

type UserProfileHandler struct {
	userRepository repository.Repository
	bh.BaseHandler
}

func NewUserProfileHandler(log *logrus.Logger, rep repository.Repository) *UserProfileHandler {
	h := &UserProfileHandler{
		BaseHandler:    *bh.NewBaseHandler(log),
		userRepository: rep,
	}
	h.AddMethod(http.MethodGet, h.GET)
	h.AddMethod(http.MethodPost, h.POST)
	return h
}

func (h *UserProfileHandler) GET(ctx echo.Context) error {
	nickname, status := h.GetStringFromParam(ctx, "nickname")
	if status == bh.EmptyQuery {
		ctx.Response().WriteHeader(http.StatusBadRequest)
		return nil
	}

	u, err := h.userRepository.Get(nickname)
	if err != nil {
		h.UsecaseError(ctx, err, codesByErrorsGET)
		return nil
	}

	//h.Log(ctx).Debugf("get user %v", u)
	h.Respond(ctx, http.StatusOK, http_delivery.ToUserResponse(u))
	return nil
}

func (h *UserProfileHandler) POST(ctx echo.Context) error {
	req := &http_delivery.UserUpdateRequest{}
	err := h.GetRequestBody(ctx, req)
	if err != nil {
		//h.Log(ctx).Warnf("can not parse request %s", err)
		h.Error(ctx, http.StatusUnprocessableEntity, handler_errors.InvalidBody)
		return nil
	}

	nickname, status := h.GetStringFromParam(ctx, "nickname")
	if status == bh.EmptyQuery {
		ctx.Response().WriteHeader(http.StatusBadRequest)
		return nil
	}

	u, err := h.userRepository.Update(&user.User{
		Nickname: nickname,
		About:    req.About,
		Email:    req.Email,
		Fullname: req.Fullname,
	})

	if err != nil {
		h.UsecaseError(ctx, err, codesByErrorsPOST)
		return nil
	}

	//h.Log(ctx).Debugf("update user %v", u)
	h.Respond(ctx, http.StatusOK, http_delivery.ToUserResponse(u))
	return nil
}
