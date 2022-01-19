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
type ThreadsResponse []ThreadResponse

func ToThreadsResponse(thrs []forum.Thread) *ThreadsResponse {
	res := ThreadsResponse{}
	for _, thr := range thrs {
		res = append(res, *ToThreadResponse(&thr))
	}
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
type UsersResponse []UserResponse

func ToUsersResponse(usrs []forum.User) *UsersResponse {
	res := UsersResponse{}
	for _, usr := range usrs {
		res = append(res, *ToUserResponse(&usr))
	}
	return &res
}
