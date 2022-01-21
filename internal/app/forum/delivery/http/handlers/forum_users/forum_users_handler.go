package forum_users_handler

import (
	"github.com/labstack/echo/v4"
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

func (h *ForumUsersHandler) GET(ctx echo.Context) error {
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

	usrs, err := h.forumRepository.GetUsers(slug, &forum.PaginationUser{
		Limit: pag.Limit,
		Desc:  pag.Desc,
		Since: pag.Since,
	})

	if err != nil {
		h.UsecaseError(ctx, err, codesByErrorsGET)
		return nil
	}

	//h.Log(ctx).Debugf("get usrs %v", usrs)
	h.Respond(ctx, http.StatusOK, http_delivery.ToUsersResponse(usrs))
	return nil
}
