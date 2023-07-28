package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title       string      `json:"title" gorm:"type:varchar(255)"`
	Description string      `json:"description" gorm:"type:varchar(255)"`
	CreatedBy   int         `json:"created_by"`
	User        User        `json:"user" gorm:"foreignKey:CreatedBy"`
	PostImage   []PostImage `json:"post_image" gorm:"foreignKey:PostID"`
}

type PostResponse struct {
	gorm.Model
	Title       string      `json:"title" gorm:"type:varchar(255)"`
	Description string      `json:"description" gorm:"type:varchar(255)"`
	CreatedBy   int         `json:"created_by"`
	User        User        `json:"user" gorm:"foreignKey:CreatedBy"`
	PostImage   []PostImage `json:"post_image"`
}

func (PostResponse) Tablename() string {
	return "post"
}
