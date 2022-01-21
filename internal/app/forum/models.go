package forum

import "time"

//go:generate easyjson -all -disallow_unknown_fields models.go

type Forum struct {
	Title   string `json:"title"`
	User    string `json:"user"`
	Slug    string `json:"slug"`
	Posts   int64  `json:"posts"`
	Threads int64  `json:"threads"`
}

type User struct {
	Nickname string `json:"nickname"`
	Fullname string `json:"fullname"`
	About    string `json:"about"`
	Email    string `json:"email"`
}

type Thread struct {
	Id      int64     `json:"id"`
	Title   string    `json:"title"`
	Author  string    `json:"author"`
	Forum   string    `json:"forum"`
	Message string    `json:"message"`
	Votes   int64     `json:"votes"`
	Slug    string    `json:"slug"`
	Created time.Time `json:"created"`
}

type PaginationUser struct {
	Limit int64
	Desc  bool
	Since string
}

type PaginationThread struct {
	Limit int64
	Desc  bool
	Since *time.Time
}
