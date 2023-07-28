package repositories

import (
	"juna/models"

	"gorm.io/gorm"
)

type PostRepository interface {
	GetPost(ID int) (models.Post, error)
	FindPosts() ([]models.Post, error)
	CreatePost(post models.Post) (models.Post, error)
}

func RepositoryPost(db *gorm.DB) *repository {
	return &repository{db}
}

// get post
func (r *repository) GetPost(ID int) (models.Post, error) {
	var post models.Post
	err := r.db.Debug().Preload("PostImage").Preload("User").First(&post, ID).Error

	return post, err
}

// find post
func (r *repository) FindPosts() ([]models.Post, error) {
	var posts []models.Post
	err := r.db.Preload("PostImage").Preload("User").Find(&posts).Error

	return posts, err
}

// create post
func (r *repository) CreatePost(post models.Post) (models.Post, error) {
	err := r.db.Create(&post).Error

	return post, err
}
