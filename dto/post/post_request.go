package postdto

type CreatePostRequest struct {
	Title       string `json:"title" gorm:"type: varchar(255)" validate:"required"`
	Description string `json:"description" gorm:"type: varchar(255)" validate:"required"`
}
