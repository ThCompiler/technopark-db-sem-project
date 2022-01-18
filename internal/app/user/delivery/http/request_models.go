package http_delivery

//go:generate easyjson -all -disallow_unknown_fields request_models.go

//easyjson:json
type UserUpdateRequest struct {
	Fullname string `json:"fullname"`
	About    string `json:"about,omitempty"`
	Email    string `json:"email"`
}