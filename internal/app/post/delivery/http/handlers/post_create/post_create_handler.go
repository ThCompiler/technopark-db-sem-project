package post_create_handler

import (
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

func (h *PostCreateHandler) POST(w http.ResponseWriter, r *http.Request) {
	req := &http_delivery.PostsCreateRequest{}
	err := h.GetRequestBody(w, r, req)
	if err != nil {
		//h.Log(w, r).Warnf("can not parse request %s", err)
		h.Error(w, r, http.StatusUnprocessableEntity, handler_errors.InvalidBody)
		return
	}

	threadPk, status := h.GetThreadSlugFromParam(w, r, "slug")
	if status == bh.EmptyQuery {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	threadId := threadPk.GetId()
	if !threadPk.IsId() {
		id, err := h.postRepository.GetThreadId(threadPk.GetSlug())
		if err != nil {
			h.Error(w, r, http.StatusNotFound, repository.NotFoundForumSlugOrUserOrThread)
			return
		}
		threadId = id
	}

	pst, err := h.postRepository.Create(http_delivery.ToPost(req), threadId)

	if err != nil {
		h.UsecaseError(w, r, err, codesByErrorsPOST)
		return
	}

	//h.Log(w, r).Debugf("create post %v", pst)
	h.Respond(w, r, http.StatusCreated, http_delivery.ToPostsResponse(pst))
	return
}
