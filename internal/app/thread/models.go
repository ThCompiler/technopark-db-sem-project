package thread

import (
	"tech-db-forum/internal/app"
	"time"
)

//go:generate easyjson -all -disallow_unknown_fields models.go

type SortType int64

const (
	Flat = SortType(iota)
	Tree
	ParentTree
)

func ConvertToSortType(tp string) SortType {
	res := Flat
	switch tp {
	case "flat":
		res = Flat
		break
	case "tree":
		res = Tree
		break
	case "parent_tree":
		res = ParentTree
		break
	}
	return res
}

type PaginationPost struct {
	Limit int64
	Desc  bool
	Since int64
	Type  SortType
}

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

type Thread struct {
	Id      int64     `json:"id"`
	Title   string    `json:"title"`
	Author  string    `json:"author"`
	Forum   string    `json:"forum"`
	Message string    `json:"message"`
	Votes   int64     `json:"votes"`
	Slug    string    `json:"slug,omitempty"`
	Created time.Time `json:"created"`
}

type ThreadPK struct {
	id   *int64
	slug *string
}

func (thr *ThreadPK) SetId(id int64) {
	thr.id = &id
	thr.slug = nil
}

func (thr *ThreadPK) SetSlug(slug string) {
	thr.slug = &slug
	thr.id = nil
}

func (thr *ThreadPK) IsId() bool {
	return thr.slug == nil
}

func (thr *ThreadPK) GetId() int64 {
	if thr.IsId() {
		return *thr.id
	}
	return app.InvalidInt
}

func (thr *ThreadPK) GetSlug() string {
	if !thr.IsId() {
		return *thr.slug
	}
	return ""
}
