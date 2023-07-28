package postdto

import "juna/models"

type DataResponse struct {
	ID          int                `json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	PostImage   []models.PostImage `json:"post_image"`
	CreatedBy   models.User        `json:"created_by"`
}

type PostResponse struct {
	Post interface{} `json:"post"`
}

type PostsResponse struct {
	Posts interface{} `json:"posts"`
}
