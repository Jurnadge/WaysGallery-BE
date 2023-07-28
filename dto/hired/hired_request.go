package hireddto

type CreateHiredRequest struct {
	Title        string `json:"title" gorm:"type: varchar(255)" validate:"required"`
	Description  string `json:"description" gorm:"type: varchar(255)" validate:"required"`
	StartProject string `json:"startProject" gorm:"type: varchar(255)" validate:"required"`
	EndProject   string `json:"endProject" gorm:"type: varchar(255)" validate:"required"`
	Price        string `json:"price" gorm:"type: varchar(255)" validate:"required"`
	OrderTo      string `json:"orderTo" gorm:"type: varchar(255)" validate:"required"`
	Status       string `json:"status" gorm:"type: varchar(255)"`
}

type UpdateHireRequest struct {
	Status string `json:"status" gorm:"type: varchar(255)" validate:"required"`
}
