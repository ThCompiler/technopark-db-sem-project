package http_delivery

import "tech-db-forum/internal/app/service"

//go:generate easyjson -all -disallow_unknown_fields response_models.go

//easyjson:json
type StatusResponse struct {
	User   string `json:"user"`
	Forum  string `json:"forum"`
	Thread string `json:"thread"`
	Post   string `json:"post"`
}

func ToStatusResponse(usr *service.Status) StatusResponse {
	return StatusResponse{
		User:   usr.User,
		Forum:  usr.Forum,
		Thread: usr.Thread,
		Post:   usr.Post,
	}
}
