package http_delivery

import (
	"tech-db-forum/internal/app/forum"
	"tech-db-forum/internal/app/post"
	"tech-db-forum/internal/app/thread"
	"tech-db-forum/internal/app/user"
)

//go:generate easyjson -all -disallow_unknown_fields response_models.go

//easyjson:json
type PostResponse struct {
	post.Post
}

func ToPostResponse(pst *post.Post) *PostResponse {
	return &PostResponse{
		Post: *pst,
	}
}

//easyjson:json
type PostsResponse []PostResponse

func ToPostsResponse(psts []post.Post) *PostsResponse {
	res := PostsResponse{}
	for _, pst := range psts {
		res = append(res, *ToPostResponse(&pst))
	}
	return &res
}

//easyjson:json
type UserResponse struct {
	user.User
}

func ToUserResponse(us *user.User) *UserResponse {
	return &UserResponse{
		User: *us,
	}
}

//easyjson:json
type ForumResponse struct {
	forum.Forum
}

func ToForumResponse(frm *forum.Forum) *ForumResponse {
	return &ForumResponse{
		Forum: *frm,
	}
}

//easyjson:json
type ThreadResponse struct {
	thread.Thread
}

func ToThreadResponse(thr *thread.Thread) *ThreadResponse {
	return &ThreadResponse{
		Thread: *thr,
	}
}

//easyjson:json
type GetPostResponse struct {
	Post   *PostResponse   `json:"post"`
	Author *UserResponse   `json:"author,omitempty"`
	Thread *ThreadResponse `json:"thread,omitempty"`
	Forum  *ForumResponse  `json:"forum,omitempty"`
}
