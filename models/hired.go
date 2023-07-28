package models

import (
	"time"

	"gorm.io/gorm"
)

type Hired struct {
	gorm.Model
	Title        string    `json:"title" gorm:"type: varchar(255)"`
	Description  string    `json:"description" gorm:"type: text"`
	StartProject time.Time `json:"startProject"`
	EndProject   time.Time `json:"endProject"`
	Price        int       `json:"price"`
	Status       string    `json:"status" gorm:"type: varchar(255)"`
	OrderTo      int       `json:"-"`
	OrderBy      int       `json:"-"`
	UserOrderTo  User      `json:"orderTo" gorm:"foreignKey:OrderTo"`
	UserOrderBy  User      `json:"orderBy" gorm:"foreignKey:OrderBy"`
}
