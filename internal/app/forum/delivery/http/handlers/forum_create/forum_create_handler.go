package forum_create_handler

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"tech-db-forum/internal/app/forum"
	"tech-db-forum/internal/app/forum/delivery/http"
	"tech-db-forum/internal/app/forum/repository"
	bh "tech-db-forum/internal/pkg/handler"
	"tech-db-forum/internal/pkg/handler/handler_errors"
	postgresql_utilits "tech-db-forum/internal/pkg/utilits/postgresql"
)

type ForumCreateHandler struct {
	forumRepository repository.Repository
	bh.BaseHandler
}

func NewForumCreateHandler(log *logrus.Logger, rep repository.Repository) *ForumCreateHandler {
	h := &ForumCreateHandler{
		BaseHandler:    *bh.NewBaseHandler(log),
		forumRepository: rep,
	}

	h.AddMethod(http.MethodPost, h.POST)
	return h
}

func (h *ForumCreateHandler) POST(w http.ResponseWriter, r *http.Request) {
	req := &http_delivery.ForumCreateRequest{}
	err := h.GetRequestBody(w, r, req)
	if err != nil {
		//h.Log(w, r).Warnf("can not parse request %s", err)
		h.Error(w, r, http.StatusUnprocessableEntity, handler_errors.InvalidBody)
		return
	}

	frm, err := h.forumRepository.Create(&forum.Forum{
		Title: req.Title,
		User: req.User,
		Slug: req.Slug,
	})

	if err == postgresql_utilits.Conflict {
		//h.Log(w, r).Warnf("conflict with request %v", req)
		h.Respond(w, r, http.StatusConflict, http_delivery.ToForumResponse(frm))
		return
	}

	if err != nil {
		h.UsecaseError(w, r, err, codesByErrorsPOST)
		return
	}

	//h.Log(w, r).Debugf("create forum %v", frm)
	h.Respond(w, r, http.StatusCreated, http_delivery.ToForumResponse(frm))
	return
}
