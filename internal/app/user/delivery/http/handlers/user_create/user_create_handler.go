package user_create_handler

import (
	"github.com/qiangxue/fasthttp-routing"
	"github.com/sirupsen/logrus"
	"net/http"
	"tech-db-forum/internal/app/user"
	"tech-db-forum/internal/app/user/delivery/http"
	"tech-db-forum/internal/app/user/repository"
	bh "tech-db-forum/internal/pkg/handler"
	"tech-db-forum/internal/pkg/handler/handler_errors"
	"tech-db-forum/internal/pkg/utilits/postgresql"
)

type UserCreateHandler struct {
	userRepository repository.Repository
	bh.BaseHandler
}

func NewUserCreateHandler(log *logrus.Logger, rep repository.Repository) *UserCreateHandler {
	h := &UserCreateHandler{
		BaseHandler:    *bh.NewBaseHandler(log),
		userRepository: rep,
	}
	h.AddMethod(http.MethodPost, h.POST)
	return h
}

func (h *UserCreateHandler) POST(ctx *routing.Context) error {
	req := &http_delivery.UserUpdateRequest{}
	err := h.GetRequestBody(ctx, req)
	if err != nil {
		h.Log(ctx).Warnf("can not parse request %s", err)
		h.Error(ctx, http.StatusUnprocessableEntity, handler_errors.InvalidBody)
		return nil
	}

	nickname, status := h.GetStringFromParam(ctx, "nickname")
	if status == bh.EmptyQuery {
		ctx.SetStatusCode(http.StatusBadRequest)
		return nil
	}

	u, err := h.userRepository.Create(&user.User{
		Nickname: nickname,
		About:    req.About,
		Email:    req.Email,
		Fullname: req.Fullname,
	})

	if err == postgresql_utilits.Conflict {
		h.Log(ctx).Warnf("conflict with request %v", req)
		h.Respond(ctx, http.StatusConflict, http_delivery.ToUsersResponse(u))
		return nil
	}

	if err != nil {
		h.UsecaseError(ctx, err, codesByErrorsPOST)
		return nil
	}

	h.Log(ctx).Debugf("create user %v", u)
	h.Respond(ctx, http.StatusCreated, http_delivery.ToUserResponse(&u[0]))
	return nil
}
