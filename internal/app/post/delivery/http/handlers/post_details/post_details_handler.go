package post_details_handler

import (
	"github.com/qiangxue/fasthttp-routing"
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

func (h *PostDetailsHandler) GET(ctx *routing.Context) error {
	id, status, err := h.GetInt64FromParam(ctx, "id")
	if err != nil {
		h.Error(ctx, status, err)
		return nil
	}

	related, _ := h.GetArrayStringFromQueries(ctx, "related")

	pst, err := h.postRepository.Get(id)
	if err != nil {
		h.UsecaseError(ctx, err, codesByErrorsGET)
		return nil
	}
	response := http_delivery.GetPostResponse{Post: http_delivery.ToPostResponse(pst)}

	if err = h.selectRelatedValueForPost(pst, &response, related); err != nil {
		h.UsecaseError(ctx, err, codesByErrorsGET)
		return nil
	}

	//h.Log(ctx).Debugf("get post %v", pst)
	h.Respond(ctx, http.StatusOK, &response)
	return nil
}

func (h *PostDetailsHandler) POST(ctx *routing.Context) error {
	req := &http_delivery.PostUpdateRequest{}
	err := h.GetRequestBody(ctx, req)
	if err != nil {
		//h.Log(ctx).Warnf("can not parse request %s", err)
		h.Error(ctx, http.StatusUnprocessableEntity, handler_errors.InvalidBody)
		return nil
	}

	id, status, err := h.GetInt64FromParam(ctx, "id")
	if err != nil {
		h.Error(ctx, status, err)
		return nil
	}

	var pst *post.Post
	if req.Message == "" {
		pst, err = h.postRepository.SetNotEdit(id)
	} else {
		pst, err = h.postRepository.Update(&post.Post{Id: id, Message: req.Message})
	}

	if err != nil {
		h.UsecaseError(ctx, err, codesByErrorsPOST)
		return nil
	}

	//h.Log(ctx).Debugf("update post %v", pst)
	h.Respond(ctx, http.StatusOK, http_delivery.ToPostResponse(pst))
	return nil
}
