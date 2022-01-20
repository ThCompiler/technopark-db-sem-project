package http_delivery

import "time"

//go:generate easyjson -all -disallow_unknown_fields request_models.go

//easyjson:json
type ThreadUpdateRequest struct {
	Title   string `json:"title"`
	Message string `json:"message,omitempty"`
}

//easyjson:json
type VoteRequest struct {
	Nickname string `json:"nickname"`
	Voice    int64  `json:"voice"`
}

//easyjson:json
type ThreadCreateRequest struct {
	Title   string    `json:"title"`
	Author  string    `json:"author"`
	Message string    `json:"message,omitempty"`
	Forum   string    `json:"forum,omitempty"`
	Slug    string    `json:"slug,omitempty"`
	Created time.Time `json:"created,omitempty"`
}
