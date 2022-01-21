package thread_details_handler

import (
	"github.com/labstack/echo/v4"
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

func (h *ThreadDetailsHandler) GET(ctx echo.Context) error {
	id, status := h.GetThreadSlugFromParam(ctx, "slug")
	if status == bh.EmptyQuery {
		h.Error(ctx, status, handler_errors.InvalidQueries)
		return nil
	}

	thr, err := h.threadRepository.Get(id)
	if err != nil {
		h.UsecaseError(ctx, err, codesByErrorsGET)
		return nil
	}

	//h.Log(ctx).Debugf("get thread %v", thr)
	h.Respond(ctx, http.StatusOK, http_delivery.ToThreadResponse(thr))
	return nil
}

func (h *ThreadDetailsHandler) POST(ctx echo.Context) error {
	req := &http_delivery.ThreadUpdateRequest{}
	err := h.GetRequestBody(ctx, req)
	if err != nil {
		//h.Log(ctx).Warnf("can not parse request %s", err)
		h.Error(ctx, http.StatusUnprocessableEntity, handler_errors.InvalidBody)
		return nil
	}

	id, status := h.GetThreadSlugFromParam(ctx, "slug")
	if status == bh.EmptyQuery {
		h.Error(ctx, status, handler_errors.InvalidQueries)
		return nil
	}

	thr, err := h.threadRepository.Update(&thread.Thread{
		Id:      id.GetId(),
		Slug:    id.GetSlug(),
		Title:   req.Title,
		Message: req.Message,
	})

	if err != nil {
		h.UsecaseError(ctx, err, codesByErrorsPOST)
		return nil
	}

	//h.Log(ctx).Debugf("update post %v", thr)
	h.Respond(ctx, http.StatusOK, http_delivery.ToThreadResponse(thr))
	return nil
}
