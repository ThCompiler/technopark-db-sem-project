package thread_create_handler

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"tech-db-forum/internal/app/thread"
	"tech-db-forum/internal/app/thread/delivery/http"
	"tech-db-forum/internal/app/thread/repository"
	bh "tech-db-forum/internal/pkg/handler"
	"tech-db-forum/internal/pkg/handler/handler_errors"
	postgresql_utilits "tech-db-forum/internal/pkg/utilits/postgresql"
)

type ThreadCreateHandler struct {
	threadRepository repository.Repository
	bh.BaseHandler
}

func NewThreadCreateHandler(log *logrus.Logger, rep repository.Repository) *ThreadCreateHandler {
	h := &ThreadCreateHandler{
		BaseHandler:      *bh.NewBaseHandler(log),
		threadRepository: rep,
	}

	h.AddMethod(http.MethodPost, h.POST)
	return h
}

func (h *ThreadCreateHandler) POST(w http.ResponseWriter, r *http.Request) {
	req := &http_delivery.ThreadCreateRequest{}
	err := h.GetRequestBody(w, r, req)
	if err != nil {
		//h.Log(w, r).Warnf("can not parse request %s", err)
		h.Error(w, r, http.StatusUnprocessableEntity, handler_errors.InvalidBody)
		return
	}

	slug, status := h.GetStringFromParam(w, r, "slug")
	if status == bh.EmptyQuery {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Forum == "" {
		req.Forum = slug
	}

	thr, err := h.threadRepository.Create(&thread.Thread{
		Slug:    req.Slug,
		Forum:   req.Forum,
		Title:   req.Title,
		Message: req.Message,
		Created: req.Created,
		Author:  req.Author,
	})

	if err == postgresql_utilits.Conflict {
		//h.Log(w, r).Warnf("conflict with request %v", req)
		h.Respond(w, r, http.StatusConflict, http_delivery.ToThreadResponse(thr))
		return
	}

	if err != nil {
		h.UsecaseError(w, r, err, codesByErrorsPOST)
		return
	}

	//h.Log(w, r).Debugf("create thread %v", thr)
	h.Respond(w, r, http.StatusCreated, http_delivery.ToThreadResponse(thr))
	return
}
