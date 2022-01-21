package thread_vote_handler

import (
	routing "github.com/qiangxue/fasthttp-routing"
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

func (h *ThreadVoteHandler) POST(ctx *routing.Context) error {
	req := &http_delivery.VoteRequest{}
	err := h.GetRequestBody(ctx, req)
	if err != nil {
		//h.Log(ctx).Warnf("can not parse request %s", err)
		h.Error(ctx, http.StatusUnprocessableEntity, handler_errors.InvalidBody)
		return nil
	}

	id, status := h.GetThreadSlugFromParam(ctx, "slug")
	if status == bh.EmptyQuery {
		h.Error(ctx, status, handler_errors.InvalidQueries)
		return nil
	}

	thr, err := h.threadRepository.SetVote(id, req.Nickname, req.Voice)

	if err != nil {
		h.UsecaseError(ctx, err, codesByErrorsGET)
		return nil
	}

	//h.Log(ctx).Debugf("set vote %v", thr)
	h.Respond(ctx, http.StatusOK, http_delivery.ToThreadResponse(thr))
	return nil
}
