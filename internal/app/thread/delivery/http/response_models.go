package http_delivery

import (
	"tech-db-forum/internal/app/thread"
)

//go:generate easyjson -all -disallow_unknown_fields response_models.go

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
type PostResponse struct {
	thread.Post
}

func ToPostResponse(pst *thread.Post) *PostResponse {
	return &PostResponse{
		Post: *pst,
	}
}

//easyjson:json
type PostsResponse []PostResponse

func ToPostsResponse(psts []thread.Post) *PostsResponse {
	res := PostsResponse{}
	for _, pst := range psts {
		res = append(res, *ToPostResponse(&pst))
	}
	return &res
}
