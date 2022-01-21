package forum_details_handler

import (
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

func (h *ForumDetailsHandler) GET(w http.ResponseWriter, r *http.Request) {
	id, status := h.GetStringFromParam(w, r, "slug")
	if status == bh.EmptyQuery {
		h.Error(w, r, http.StatusBadRequest, handler_errors.InvalidQueries)
		return
	}

	frm, err := h.forumRepository.Get(id)
	if err != nil {
		h.UsecaseError(w, r, err, codesByErrorsGET)
		return
	}

	//h.Log(w, r).Debugf("get post %v", frm)
	h.Respond(w, r, http.StatusOK, http_delivery.ToForumResponse(frm))
	return
}
