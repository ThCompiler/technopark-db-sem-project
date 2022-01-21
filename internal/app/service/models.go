package service

//go:generate easyjson -all -disallow_unknown_fields models.go

type Status struct {
	User int32 `json:"user"`
	Forum int32 `json:"forum"`
	Thread int32 `json:"thread"`
	Post int64 `json:"post"`
}