package http_delivery

import (
	"tech-db-forum/internal/app/forum"
)

//go:generate easyjson -all -disallow_unknown_fields response_models.go

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
	forum.Thread
}

func ToThreadResponse(thr *forum.Thread) *ThreadResponse {
	return &ThreadResponse{
		Thread: *thr,
	}
}

//easyjson:json
type ThreadsResponse []forum.Thread

func ToThreadsResponse(thrs []forum.Thread) *ThreadsResponse {
	if thrs == nil {
		thrs = []forum.Thread{}
	}
	res := ThreadsResponse(thrs)
	return &res
}

//easyjson:json
type UserResponse struct {
	forum.User
}

func ToUserResponse(usr *forum.User) *UserResponse {
	return &UserResponse{
		User: *usr,
	}
}

//easyjson:json
type UsersResponse []forum.User

func ToUsersResponse(usrs []forum.User) *UsersResponse {
	if usrs == nil {
		usrs = []forum.User{}
	}
	res := UsersResponse(usrs)
	return &res
}
