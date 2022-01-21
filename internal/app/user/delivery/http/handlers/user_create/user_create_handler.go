package user_create_handler

import (
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

func (h *UserCreateHandler) POST(w http.ResponseWriter, r *http.Request) {
	req := &http_delivery.UserUpdateRequest{}
	err := h.GetRequestBody(w, r, req)
	if err != nil {
		//h.Log(w, r).Warnf("can not parse request %s", err)
		h.Error(w, r, http.StatusUnprocessableEntity, handler_errors.InvalidBody)
		return
	}

	nickname, status := h.GetStringFromParam(w, r, "nickname")
	if status == bh.EmptyQuery {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := h.userRepository.Create(&user.User{
		Nickname: nickname,
		About:    req.About,
		Email:    req.Email,
		Fullname: req.Fullname,
	})

	if err == postgresql_utilits.Conflict {
		//h.Log(w, r).Warnf("conflict with request %v", req)
		h.Respond(w, r, http.StatusConflict, http_delivery.ToUsersResponse(u))
		return
	}

	if err != nil {
		h.UsecaseError(w, r, err, codesByErrorsPOST)
		return
	}

	//h.Log(w, r).Debugf("create user %v", u)
	h.Respond(w, r, http.StatusCreated, http_delivery.ToUserResponse(&u[0]))
	return
}
