package http_delivery

import "tech-db-forum/internal/app/post"

//go:generate easyjson -all -disallow_unknown_fields request_models.go

//easyjson:json
type PostUpdateRequest struct {
	Message string `json:"message,omitempty"`
}

//easyjson:json
type PostCreateRequest struct {
	Parent  int64  `json:"parent"`
	Author  string `json:"author"`
	Message string `json:"message,omitempty"`
}

//easyjson:json
type PostsCreateRequest []PostCreateRequest

func ToPost(psts *PostsCreateRequest) []post.Post {
	var res []post.Post
	for _, pst := range *psts {
		res = append(res, post.Post{Parent: pst.Parent, Author: pst.Author, Message: pst.Message})
	}

	return res
}
