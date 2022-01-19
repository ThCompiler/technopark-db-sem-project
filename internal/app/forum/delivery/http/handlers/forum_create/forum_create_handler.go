package forum_create_handler

import (
	routing "github.com/qiangxue/fasthttp-routing"
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

func (h *ForumCreateHandler) POST(ctx *routing.Context) error {
	req := &http_delivery.ForumCreateRequest{}
	err := h.GetRequestBody(ctx, req)
	if err != nil {
		h.Log(ctx).Warnf("can not parse request %s", err)
		h.Error(ctx, http.StatusUnprocessableEntity, handler_errors.InvalidBody)
		return nil
	}

	frm, err := h.forumRepository.Create(&forum.Forum{
		Title: req.Title,
		User: req.User,
		Slug: req.Slug,
	})

	if err == postgresql_utilits.Conflict {
		h.Log(ctx).Warnf("conflict with request %v", req)
		h.Respond(ctx, http.StatusConflict, http_delivery.ToForumResponse(frm))
		return nil
	}

	if err != nil {
		h.UsecaseError(ctx, err, codesByErrorsPOST)
		return nil
	}

	h.Log(ctx).Debugf("create forum %v", frm)
	h.Respond(ctx, http.StatusCreated, http_delivery.ToForumResponse(frm))
	return nil
}
