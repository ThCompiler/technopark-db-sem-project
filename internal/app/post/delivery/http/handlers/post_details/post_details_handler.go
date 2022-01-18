package post_details_handler

import (
	"github.com/qiangxue/fasthttp-routing"
	"github.com/sirupsen/logrus"
	"net/http"
	"sync"
	repForum "tech-db-forum/internal/app/forum/repository"
	"tech-db-forum/internal/app/post"
	"tech-db-forum/internal/app/post/delivery/http"
	"tech-db-forum/internal/app/post/repository"
	"tech-db-forum/internal/app/thread"
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
	h.AddMethod(http.MethodPost, h.POST)
	return h
}

func (h *PostDetailsHandler) GET(ctx *routing.Context) error {
	id, status, err := h.GetInt64FromParam(ctx, "id")
	if status == bh.EmptyQuery {
		h.Error(ctx, status, err)
		return nil
	}

	related, status := h.GetArrayStringFromQueries(ctx, "related")
	if status == bh.EmptyQuery {
		h.Error(ctx, status, handler_errors.InvalidQueries)
		return nil
	}

	pst, err := h.postRepository.Get(id)
	if err != nil {
		h.UsecaseError(ctx, err, codesByErrorsGET)
		return nil
	}
	response := http_delivery.GetPostResponse{Post: http_delivery.ToPostResponse(pst)}

	if err = h.selectRelatedValueForPost(pst, &response, &related); err != nil {
		h.UsecaseError(ctx, err, codesByErrorsGET)
		return nil
	}

	h.Log(ctx).Debugf("get post %v", pst)
	h.Respond(ctx, http.StatusOK, &response)
	return nil
}

func (h *PostDetailsHandler) POST(ctx *routing.Context) error {
	req := &http_delivery.PostUpdateRequest{}
	err := h.GetRequestBody(ctx, req)
	if err != nil {
		h.Log(ctx).Warnf("can not parse request %s", err)
		h.Error(ctx, http.StatusUnprocessableEntity, handler_errors.InvalidBody)
		return nil
	}

	id, status, err := h.GetInt64FromParam(ctx, "id")
	if status == bh.EmptyQuery {
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

	h.Log(ctx).Debugf("update post %v", pst)
	h.Respond(ctx, http.StatusCreated, http_delivery.ToPostResponse(pst))
	return nil
}

func (h *PostDetailsHandler) getUserInfo(usr *http_delivery.UserResponse,
	nickname string, wait *sync.WaitGroup, resErr *error) {
	defer wait.Done()
	if usr == nil {
		return
	}

	res, err := h.userRepository.Get(nickname)
	if err != nil {
		resErr = &err
		return
	}
	usr = http_delivery.ToUserResponse(res)
}

func (h *PostDetailsHandler) getForumInfo(frm *http_delivery.ForumResponse,
	slug string, wait *sync.WaitGroup, resErr *error) {
	defer wait.Done()
	if frm == nil {
		return
	}

	res, err := h.forumRepository.Get(slug)
	if err != nil {
		resErr = &err
		return
	}
	frm = http_delivery.ToForumResponse(res)
}

func (h *PostDetailsHandler) getThreadInfo(thr *http_delivery.ThreadResponse, id int64,
	wait *sync.WaitGroup, resErr *error) {
	defer wait.Done()
	if thr == nil {
		return
	}

	threadPk := &thread.ThreadPK{}
	threadPk.SetId(id)
	res, err := h.threadRepository.Get(threadPk)
	if err != nil {
		resErr = &err
		return
	}
	thr = http_delivery.ToThreadResponse(res)
}

func (h *PostDetailsHandler) selectRelatedValueForPost(pst *post.Post,
	ans *http_delivery.GetPostResponse, related *[]string) error {

	var UserErr, ThreadErr, ForumErr *error

	ThreadRes := &http_delivery.ThreadResponse{}
	ThreadRes = nil
	ForumRes := &http_delivery.ForumResponse{}
	ForumRes = nil
	UserRes := &http_delivery.UserResponse{}
	UserRes = nil

	wait := &sync.WaitGroup{}
	for _, value := range *related {
		wait.Add(1)
		switch value {
		case "user":
			go h.getUserInfo(UserRes, pst.Author, wait, UserErr)
			break
		case "forum":
			go h.getForumInfo(ForumRes, pst.Forum, wait, ForumErr)
			break
		case "thread":
			go h.getThreadInfo(ThreadRes, pst.Thread, wait, ThreadErr)
			break
		}
	}
	wait.Wait()

	if UserErr != nil {
		return *UserErr
	}

	if ForumErr != nil {
		return *ForumErr
	}

	if ThreadErr != nil {
		return *ThreadErr
	}

	if ThreadRes != nil {
		ans.Thread = *ThreadRes
	}

	if ForumRes != nil {
		ans.Forum = *ForumRes
	}

	if UserRes != nil {
		ans.Author = *UserRes
	}
	return nil
}
