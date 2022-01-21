package thread_details_handler

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"tech-db-forum/internal/app/thread"
	"tech-db-forum/internal/app/thread/delivery/http"
	"tech-db-forum/internal/app/thread/repository"
	bh "tech-db-forum/internal/pkg/handler"
	"tech-db-forum/internal/pkg/handler/handler_errors"
)

type ThreadDetailsHandler struct {
	threadRepository repository.Repository
	bh.BaseHandler
}

func NewThreadDetailsHandler(log *logrus.Logger, rep repository.Repository) *ThreadDetailsHandler {
	h := &ThreadDetailsHandler{
		BaseHandler:      *bh.NewBaseHandler(log),
		threadRepository: rep,
	}
	h.AddMethod(http.MethodGet, h.GET)
	h.AddMethod(http.MethodPost, h.POST)
	return h
}

func (h *ThreadDetailsHandler) GET(w http.ResponseWriter, r *http.Request) {
	id, status := h.GetThreadSlugFromParam(w, r, "slug")
	if status == bh.EmptyQuery {
		h.Error(w, r, status, handler_errors.InvalidQueries)
		return
	}

	thr, err := h.threadRepository.Get(id)
	if err != nil {
		h.UsecaseError(w, r, err, codesByErrorsGET)
		return
	}

	//h.Log(w, r).Debugf("get thread %v", thr)
	h.Respond(w, r, http.StatusOK, http_delivery.ToThreadResponse(thr))
	return
}

func (h *ThreadDetailsHandler) POST(w http.ResponseWriter, r *http.Request) {
	req := &http_delivery.ThreadUpdateRequest{}
	err := h.GetRequestBody(w, r, req)
	if err != nil {
		//h.Log(w, r).Warnf("can not parse request %s", err)
		h.Error(w, r, http.StatusUnprocessableEntity, handler_errors.InvalidBody)
		return
	}

	id, status := h.GetThreadSlugFromParam(w, r, "slug")
	if status == bh.EmptyQuery {
		h.Error(w, r, status, handler_errors.InvalidQueries)
		return
	}

	thr, err := h.threadRepository.Update(&thread.Thread{
		Id:      id.GetId(),
		Slug:    id.GetSlug(),
		Title:   req.Title,
		Message: req.Message,
	})

	if err != nil {
		h.UsecaseError(w, r, err, codesByErrorsPOST)
		return
	}

	//h.Log(w, r).Debugf("update post %v", thr)
	h.Respond(w, r, http.StatusOK, http_delivery.ToThreadResponse(thr))
	return
}
