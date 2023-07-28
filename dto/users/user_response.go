package usersdto

import "juna/models"

type UpdateResponse struct {
	Avatar   string `json:"avatar"`
	Greeting string `json:"greeting"`
	Fullname string `json:"fullname"`
}

type UserDetailResponse struct {
	ID       int           `json:"id"`
	Fullname string        `json:"fullname"`
	Email    string        `json:"email"`
	Avatar   string        `json:"avatar"`
	Greeting string        `json:"greeting"`
	Post     []models.Post `json:"posts"`
	Arts     string        `json:"arts"`
}

type UserPostResponse struct {
	ID          int                `json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	PostImage   []models.PostImage `json:"post_image"`
}

type UserArtResponse struct {
	ID    string `json:"id"`
	Image string `json:"image"`
}

type UserRespnse struct {
	User interface{} `json:"user"`
}
