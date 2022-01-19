package http_delivery

import "tech-db-forum/internal/app/user"

//go:generate easyjson -all -disallow_unknown_fields response_models.go

//easyjson:json
type UserResponse struct {
	Nickname string `json:"nickname"`
	Fullname string `json:"fullname"`
	About    string `json:"about,omitempty"`
	Email    string `json:"email"`
}

func ToUserResponse(usr *user.User) *UserResponse {
	return &UserResponse{
		Nickname: usr.Nickname,
		Fullname: usr.Fullname,
		About:    usr.About,
		Email:    usr.Email,
	}
}

//easyjson:json
type UsersResponse []UserResponse

func ToUsersResponse(usrs []user.User) *UsersResponse {
	res := UsersResponse{}
	for _, usr := range usrs {
		res = append(res, *ToUserResponse(&usr))
	}
	return &res
}
