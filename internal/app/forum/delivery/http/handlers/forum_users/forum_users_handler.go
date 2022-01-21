package forum_users_handler

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"tech-db-forum/internal/app/forum"
	"tech-db-forum/internal/app/forum/delivery/http"
	"tech-db-forum/internal/app/forum/repository"
	bh "tech-db-forum/internal/pkg/handler"
	"tech-db-forum/internal/pkg/handler/handler_errors"
)

type ForumUsersHandler struct {
	forumRepository repository.Repository
	bh.BaseHandler
}

func NewForumUsersHandler(log *logrus.Logger, rep repository.Repository) *ForumUsersHandler {
	h := &ForumUsersHandler{
		BaseHandler:     *bh.NewBaseHandler(log),
		forumRepository: rep,
	}
	h.AddMethod(http.MethodGet, h.GET)
	return h
}

func (h *ForumUsersHandler) GET(w http.ResponseWriter, r *http.Request) {
	slug, status := h.GetStringFromParam(w, r, "slug")
	if status == bh.EmptyQuery {
		h.Error(w, r, http.StatusBadRequest, handler_errors.InvalidQueries)
		return
	}

	pag, status, err := h.GetPaginationFromQuery(w, r)
	if err != nil {
		h.Error(w, r, status, err)
		return
	}

	usrs, err := h.forumRepository.GetUsers(slug, &forum.PaginationUser{
		Limit: pag.Limit,
		Desc:  pag.Desc,
		Since: pag.Since,
	})

	if err != nil {
		h.UsecaseError(w, r, err, codesByErrorsGET)
		return
	}

	//h.Log(w, r).Debugf("get usrs %v", usrs)
	h.Respond(w, r, http.StatusOK, http_delivery.ToUsersResponse(usrs))
	return
}
