package post

import "time"

//go:generate easyjson -all -disallow_unknown_fields models.go

type Post struct {
	Id       int64     `json:"id"`
	Parent   int64     `json:"parent"`
	Author   string    `json:"author"`
	Message  string    `json:"message"`
	Is_Edited bool      `json:"isEdited"`
	Forum    string    `json:"forum"`
	Thread   int64     `json:"thread"`
	Created  time.Time `json:"created"`
}
