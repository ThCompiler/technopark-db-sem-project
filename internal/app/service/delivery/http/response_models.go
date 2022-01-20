package http_delivery

import "tech-db-forum/internal/app/service"

//go:generate easyjson -all -disallow_unknown_fields response_models.go

//easyjson:json
type StatusResponse struct {
	User   int32 `json:"user"`
	Forum  int32 `json:"forum"`
	Thread int32 `json:"thread"`
	Post   int64 `json:"post"`
}

func ToStatusResponse(usr *service.Status) *StatusResponse {
	return &StatusResponse{
		User:   usr.User,
		Forum:  usr.Forum,
		Thread: usr.Thread,
		Post:   usr.Post,
	}
}
