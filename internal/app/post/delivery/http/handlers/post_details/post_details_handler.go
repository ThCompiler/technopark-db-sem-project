package post_details_handler

import (
	"github.com/sirupsen/logrus"
	"net/http"
	repForum "tech-db-forum/internal/app/forum/repository"
	"tech-db-forum/internal/app/post"
	"tech-db-forum/internal/app/post/delivery/http"
	"tech-db-forum/internal/app/post/repository"
	repThread "tech-db-forum/internal/app/thread/repository"
	repUser "tech-db-forum/internal/app/user/repository"
	bh "tech-db-forum/internal/pkg/handler"
	"tech-db-forum/internal/pkg/handler/handler_errors"
)

type PostDetailsHandler struct {
	postRepository   repository.Repository
	threadRepository repThread.Repository
	userRepository   repUser.Repository
	forumRepository  repForum.Repository
	bh.BaseHandler
}

func NewPostDetailsHandler(log *logrus.Logger, rep repository.Repository, repTh repThread.Repository,
	repUs repUser.Repository, repFr repForum.Repository) *PostDetailsHandler {
	h := &PostDetailsHandler{
		BaseHandler:      *bh.NewBaseHandler(log),
		postRepository:   rep,
		threadRepository: repTh,
		userRepository:   repUs,
		forumRepository:  repFr,
	}
	h.AddMethod(http.MethodGet, h.GET)
	h.AddMethod(http.MethodPost, h.POST)
	return h
}

func (h *PostDetailsHandler) GET(w http.ResponseWriter, r *http.Request) {
	id, status, err := h.GetInt64FromParam(w, r, "id")
	if err != nil {
		h.Error(w, r, status, err)
		return
	}

	related, _ := h.GetArrayStringFromQueries(w, r, "related")

	pst, err := h.postRepository.Get(id)
	if err != nil {
		h.UsecaseError(w, r, err, codesByErrorsGET)
		return
	}
	response := http_delivery.GetPostResponse{Post: http_delivery.ToPostResponse(pst)}

	if err = h.selectRelatedValueForPost(pst, &response, related); err != nil {
		h.UsecaseError(w, r, err, codesByErrorsGET)
		return
	}

	//h.Log(w, r).Debugf("get post %v", pst)
	h.Respond(w, r, http.StatusOK, &response)
	return
}

func (h *PostDetailsHandler) POST(w http.ResponseWriter, r *http.Request) {
	req := &http_delivery.PostUpdateRequest{}
	err := h.GetRequestBody(w, r, req)
	if err != nil {
		//h.Log(w, r).Warnf("can not parse request %s", err)
		h.Error(w, r, http.StatusUnprocessableEntity, handler_errors.InvalidBody)
		return
	}

	id, status, err := h.GetInt64FromParam(w, r, "id")
	if err != nil {
		h.Error(w, r, status, err)
		return
	}

	var pst *post.Post
	if req.Message == "" {
		pst, err = h.postRepository.SetNotEdit(id)
	} else {
		pst, err = h.postRepository.Update(&post.Post{Id: id, Message: req.Message})
	}

	if err != nil {
		h.UsecaseError(w, r, err, codesByErrorsPOST)
		return
	}

	//h.Log(w, r).Debugf("update post %v", pst)
	h.Respond(w, r, http.StatusOK, http_delivery.ToPostResponse(pst))
	return
}
