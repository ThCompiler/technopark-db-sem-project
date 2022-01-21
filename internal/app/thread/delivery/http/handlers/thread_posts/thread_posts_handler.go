package thread_posts_handler

import (
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"tech-db-forum/internal/app"
	"tech-db-forum/internal/app/thread"
	"tech-db-forum/internal/app/thread/delivery/http"
	"tech-db-forum/internal/app/thread/repository"
	bh "tech-db-forum/internal/pkg/handler"
	"tech-db-forum/internal/pkg/handler/handler_errors"
)

type ThreadPostsHandler struct {
	threadRepository repository.Repository
	bh.BaseHandler
}

func NewThreadPostsHandler(log *logrus.Logger, rep repository.Repository) *ThreadPostsHandler {
	h := &ThreadPostsHandler{
		BaseHandler:    *bh.NewBaseHandler(log),
		threadRepository: rep,
	}

	h.AddMethod(http.MethodGet, h.GET)
	return h
}

func (h *ThreadPostsHandler) GET(ctx *routing.Context) error {
	id, status := h.GetThreadSlugFromParam(ctx, "slug")
	if status == bh.EmptyQuery {
		h.Error(ctx, status, handler_errors.InvalidQueries)
		return nil
	}

	pag, status, err := h.GetPaginationFromQuery(ctx)
	if err != nil {
		h.Error(ctx, status, err)
		return nil
	}

	since, err := strconv.ParseInt(pag.Since, 10, 64)
	if err != nil && pag.Since != "" {
		h.Error(ctx, status, handler_errors.InvalidQueries)
		return nil
	}

	if pag.Since == "" {
		since = app.InvalidInt
	}

	sortType, status := h.GetStringFromQueries(ctx, "sort")
	if status == bh.EmptyQuery {
		sortType = "flat"
	}

	psts, err := h.threadRepository.GetPosts(id, &thread.PaginationPost{
		Desc: pag.Desc,
		Limit: pag.Limit,
		Since: since,
		Type: thread.ConvertToSortType(sortType),
	})

	if err != nil {
		h.UsecaseError(ctx, err, codesByErrorsGET)
		return nil
	}

	//h.Log(ctx).Debugf("get psts %v", psts)
	h.Respond(ctx, http.StatusOK, http_delivery.ToPostsResponse(psts))
	return nil
}
