package http_delivery

//go:generate easyjson -all -disallow_unknown_fields request_models.go

//easyjson:json
type ForumCreateRequest struct {
	Title string  `json:"title"`
	User  string `json:"user"`
	Slug  string `json:"slug"`
}
