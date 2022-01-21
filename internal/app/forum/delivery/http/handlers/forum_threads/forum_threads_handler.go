package forum_threads_handler

import (
	"github.com/qiangxue/fasthttp-routing"
	"github.com/sirupsen/logrus"
	"net/http"
	"tech-db-forum/internal/app/forum"
	"tech-db-forum/internal/app/forum/delivery/http"
	"tech-db-forum/internal/app/forum/repository"
	bh "tech-db-forum/internal/pkg/handler"
	"tech-db-forum/internal/pkg/handler/handler_errors"
	"time"
)

type ForumThreadsHandler struct {
	forumRepository repository.Repository
	bh.BaseHandler
}

func NewForumThreadsHandler(log *logrus.Logger, rep repository.Repository) *ForumThreadsHandler {
	h := &ForumThreadsHandler{
		BaseHandler:     *bh.NewBaseHandler(log),
		forumRepository: rep,
	}
	h.AddMethod(http.MethodGet, h.GET)
	return h
}

func (h *ForumThreadsHandler) GET(ctx *routing.Context) error {
	slug, status := h.GetStringFromParam(ctx, "slug")
	if status == bh.EmptyQuery {
		h.Error(ctx, http.StatusBadRequest, handler_errors.InvalidQueries)
		return nil
	}

	pag, status, err := h.GetPaginationFromQuery(ctx)
	if err != nil {
		h.Error(ctx, status, err)
		return nil
	}

	date, err := time.Parse(time.RFC3339, pag.Since)
	if err != nil && pag.Since != "" {
		h.Error(ctx, http.StatusBadRequest, handler_errors.InvalidQueries)
		return nil
	}

	since := (*time.Time)(nil)
	if  pag.Since != "" {
		since = &date
	}

	thrs, err := h.forumRepository.GetThreads(slug, &forum.PaginationThread{Limit: pag.Limit, Desc: pag.Desc, Since: since})
	if err != nil {
		h.UsecaseError(ctx, err, codesByErrorsGET)
		return nil
	}

	//h.Log(ctx).Debugf("get threads %v", thrs)
	h.Respond(ctx, http.StatusOK, http_delivery.ToThreadsResponse(thrs))
	return nil
}
