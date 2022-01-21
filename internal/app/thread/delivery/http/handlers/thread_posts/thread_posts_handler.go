package thread_posts_handler

import (
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

func (h *ThreadPostsHandler) GET(w http.ResponseWriter, r *http.Request) {
	id, status := h.GetThreadSlugFromParam(w, r, "slug")
	if status == bh.EmptyQuery {
		h.Error(w, r, status, handler_errors.InvalidQueries)
		return
	}

	pag, status, err := h.GetPaginationFromQuery(w, r)
	if err != nil {
		h.Error(w, r, status, err)
		return
	}

	since, err := strconv.ParseInt(pag.Since, 10, 64)
	if err != nil && pag.Since != "" {
		h.Error(w, r, status, handler_errors.InvalidQueries)
		return
	}

	if pag.Since == "" {
		since = app.InvalidInt
	}

	sortType, status := h.GetStringFromQueries(w, r, "sort")
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
		h.UsecaseError(w, r, err, codesByErrorsGET)
		return
	}

	//h.Log(w, r).Debugf("get psts %v", psts)
	h.Respond(w, r, http.StatusOK, http_delivery.ToPostsResponse(psts))
	return
}
