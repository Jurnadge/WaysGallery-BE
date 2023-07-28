package database

import (
	"fmt"
	"juna/models"
	"juna/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Art{},
		&models.Post{},
		&models.PostImage{},
		&models.Hired{},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Migration Success")
}
