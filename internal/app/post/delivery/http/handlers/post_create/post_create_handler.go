package post_create_handler

import (
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/sirupsen/logrus"
	"net/http"
	"tech-db-forum/internal/app/post/delivery/http"
	"tech-db-forum/internal/app/post/repository"
	bh "tech-db-forum/internal/pkg/handler"
	"tech-db-forum/internal/pkg/handler/handler_errors"
)

type PostCreateHandler struct {
	postRepository repository.Repository
	bh.BaseHandler
}

func NewPostCreateHandler(log *logrus.Logger, rep repository.Repository) *PostCreateHandler {
	h := &PostCreateHandler{
		BaseHandler:    *bh.NewBaseHandler(log),
		postRepository: rep,
	}

	h.AddMethod(http.MethodPost, h.POST)
	return h
}

func (h *PostCreateHandler) POST(ctx *routing.Context) error {
	req := &http_delivery.PostsCreateRequest{}
	err := h.GetRequestBody(ctx, req)
	if err != nil {
		//h.Log(ctx).Warnf("can not parse request %s", err)
		h.Error(ctx, http.StatusUnprocessableEntity, handler_errors.InvalidBody)
		return nil
	}

	threadPk, status := h.GetThreadSlugFromParam(ctx, "slug")
	if status == bh.EmptyQuery {
		ctx.SetStatusCode(http.StatusBadRequest)
		return nil
	}

	threadId := threadPk.GetId()
	if !threadPk.IsId() {
		id, err := h.postRepository.GetThreadId(threadPk.GetSlug())
		if err != nil {
			h.Error(ctx, http.StatusNotFound, repository.NotFoundForumSlugOrUserOrThread)
			return nil
		}
		threadId = id
	}

	pst, err := h.postRepository.Create(http_delivery.ToPost(req), threadId)

	if err != nil {
		h.UsecaseError(ctx, err, codesByErrorsPOST)
		return nil
	}

	//h.Log(ctx).Debugf("create post %v", pst)
	h.Respond(ctx, http.StatusCreated, http_delivery.ToPostsResponse(pst))
	return nil
}
