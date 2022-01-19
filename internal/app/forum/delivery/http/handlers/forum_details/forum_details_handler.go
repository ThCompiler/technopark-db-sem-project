package forum_details_handler

import (
	"github.com/qiangxue/fasthttp-routing"
	"github.com/sirupsen/logrus"
	"net/http"
	"tech-db-forum/internal/app/forum/delivery/http"
	"tech-db-forum/internal/app/forum/repository"
	bh "tech-db-forum/internal/pkg/handler"
	"tech-db-forum/internal/pkg/handler/handler_errors"
)

type ForumDetailsHandler struct {
	forumRepository repository.Repository
	bh.BaseHandler
}

func NewForumDetailsHandler(log *logrus.Logger, rep repository.Repository) *ForumDetailsHandler {
	h := &ForumDetailsHandler{
		BaseHandler:     *bh.NewBaseHandler(log),
		forumRepository: rep,
	}
	h.AddMethod(http.MethodGet, h.GET)
	return h
}

func (h *ForumDetailsHandler) GET(ctx *routing.Context) error {
	id, status := h.GetStringFromParam(ctx, "slug")
	if status == bh.EmptyQuery {
		h.Error(ctx, http.StatusBadRequest, handler_errors.InvalidQueries)
		return nil
	}

	frm, err := h.forumRepository.Get(id)
	if err != nil {
		h.UsecaseError(ctx, err, codesByErrorsGET)
		return nil
	}

	h.Log(ctx).Debugf("get post %v", frm)
	h.Respond(ctx, http.StatusOK, http_delivery.ToForumResponse(frm))
	return nil
}
