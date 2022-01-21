package thread_create_handler

import (
	routing "github.com/qiangxue/fasthttp-routing"
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

func (h *ThreadCreateHandler) POST(ctx *routing.Context) error {
	req := &http_delivery.ThreadCreateRequest{}
	err := h.GetRequestBody(ctx, req)
	if err != nil {
		//h.Log(ctx).Warnf("can not parse request %s", err)
		h.Error(ctx, http.StatusUnprocessableEntity, handler_errors.InvalidBody)
		return nil
	}

	slug, status := h.GetStringFromParam(ctx, "slug")
	if status == bh.EmptyQuery {
		ctx.SetStatusCode(http.StatusBadRequest)
		return nil
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
		//h.Log(ctx).Warnf("conflict with request %v", req)
		h.Respond(ctx, http.StatusConflict, http_delivery.ToThreadResponse(thr))
		return nil
	}

	if err != nil {
		h.UsecaseError(ctx, err, codesByErrorsPOST)
		return nil
	}

	//h.Log(ctx).Debugf("create thread %v", thr)
	h.Respond(ctx, http.StatusCreated, http_delivery.ToThreadResponse(thr))
	return nil
}
