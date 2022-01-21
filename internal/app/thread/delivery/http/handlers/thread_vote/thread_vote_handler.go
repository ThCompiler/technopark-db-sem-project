package thread_vote_handler

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"tech-db-forum/internal/app/thread/delivery/http"
	"tech-db-forum/internal/app/thread/repository"
	bh "tech-db-forum/internal/pkg/handler"
	"tech-db-forum/internal/pkg/handler/handler_errors"
)

type ThreadVoteHandler struct {
	threadRepository repository.Repository
	bh.BaseHandler
}

func NewThreadVoteHandler(log *logrus.Logger, rep repository.Repository) *ThreadVoteHandler {
	h := &ThreadVoteHandler{
		BaseHandler:    *bh.NewBaseHandler(log),
		threadRepository: rep,
	}

	h.AddMethod(http.MethodPost, h.POST)
	return h
}

func (h *ThreadVoteHandler) POST(w http.ResponseWriter, r *http.Request) {
	req := &http_delivery.VoteRequest{}
	err := h.GetRequestBody(w, r, req)
	if err != nil {
		//h.Log(w, r).Warnf("can not parse request %s", err)
		h.Error(w, r, http.StatusUnprocessableEntity, handler_errors.InvalidBody)
		return
	}

	id, status := h.GetThreadSlugFromParam(w, r, "slug")
	if status == bh.EmptyQuery {
		h.Error(w, r, status, handler_errors.InvalidQueries)
		return
	}

	thr, err := h.threadRepository.SetVote(id, req.Nickname, req.Voice)

	if err != nil {
		h.UsecaseError(w, r, err, codesByErrorsGET)
		return
	}

	//h.Log(w, r).Debugf("set vote %v", thr)
	h.Respond(w, r, http.StatusOK, http_delivery.ToThreadResponse(thr))
	return
}
