package user_profile_handler

import (
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

func (h *UserProfileHandler) GET(w http.ResponseWriter, r *http.Request) {
	nickname, status := h.GetStringFromParam(w, r, "nickname")
	if status == bh.EmptyQuery {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := h.userRepository.Get(nickname)
	if err != nil {
		h.UsecaseError(w, r, err, codesByErrorsGET)
		return
	}

	//h.Log(w, r).Debugf("get user %v", u)
	h.Respond(w, r, http.StatusOK, http_delivery.ToUserResponse(u))
	return
}

func (h *UserProfileHandler) POST(w http.ResponseWriter, r *http.Request) {
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

	u, err := h.userRepository.Update(&user.User{
		Nickname: nickname,
		About:    req.About,
		Email:    req.Email,
		Fullname: req.Fullname,
	})

	if err != nil {
		h.UsecaseError(w, r, err, codesByErrorsPOST)
		return
	}

	//h.Log(w, r).Debugf("update user %v", u)
	h.Respond(w, r, http.StatusOK, http_delivery.ToUserResponse(u))
	return
}
