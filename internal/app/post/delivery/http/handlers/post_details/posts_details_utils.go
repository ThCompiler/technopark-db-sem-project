package post_details_handler

import (
	"sync"
	"tech-db-forum/internal/app/post"
	http_delivery "tech-db-forum/internal/app/post/delivery/http"
	"tech-db-forum/internal/app/thread"
)

type ResUserType struct {
	usr    *http_delivery.UserResponse
	status bool
}

type ResTheadType struct {
	thr    *http_delivery.ThreadResponse
	status bool
}

type ResForumType struct {
	frm    *http_delivery.ForumResponse
	status bool
}

func (h *PostDetailsHandler) getUserInfo(usr *ResUserType,
	nickname string, wait *sync.WaitGroup, resErr *error) {
	defer wait.Done()
	if usr.status {
		return
	}
	usr.status = true

	res, err := h.userRepository.Get(nickname)
	if err != nil {
		*resErr = err
		return
	}
	usr.usr = http_delivery.ToUserResponse(res)
}

func (h *PostDetailsHandler) getForumInfo(frm *ResForumType,
	slug string, wait *sync.WaitGroup, resErr *error) {
	defer wait.Done()
	if frm.status {
		return
	}
	frm.status = true

	res, err := h.forumRepository.Get(slug)
	if err != nil {
		*resErr = err
		return
	}
	frm.frm = http_delivery.ToForumResponse(res)
}

func (h *PostDetailsHandler) getThreadInfo(thr *ResTheadType, id int64,
	wait *sync.WaitGroup, resErr *error) {
	defer wait.Done()
	if thr.status {
		return
	}
	thr.status = true

	threadPk := &thread.ThreadPK{}
	threadPk.SetId(id)
	res, err := h.threadRepository.Get(threadPk)
	if err != nil {
		*resErr = err
		return
	}
	thr.thr = http_delivery.ToThreadResponse(res)
}

func (h *PostDetailsHandler) selectRelatedValueForPost(pst *post.Post,
	ans *http_delivery.GetPostResponse, related []string) error {

	var UserErr, ThreadErr, ForumErr *error

	ThreadRes := ResTheadType{
		thr:    nil,
		status: false,
	}
	ForumRes := ResForumType{
		frm:    nil,
		status: false,
	}
	UserRes := ResUserType{
		usr:    nil,
		status: false,
	}

	wait := &sync.WaitGroup{}
	for _, value := range related {
		wait.Add(1)
		switch value {
		case "user":
			go h.getUserInfo(&UserRes, pst.Author, wait, UserErr)
			break
		case "forum":
			go h.getForumInfo(&ForumRes, pst.Forum, wait, ForumErr)
			break
		case "thread":
			go h.getThreadInfo(&ThreadRes, pst.Thread, wait, ThreadErr)
			break
		default:
			wait.Done()
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

	if ThreadRes.status {
		ans.Thread = ThreadRes.thr
	}

	if ForumRes.status {
		ans.Forum = ForumRes.frm
	}

	if UserRes.status {
		ans.Author = UserRes.usr
	}
	return nil
}
