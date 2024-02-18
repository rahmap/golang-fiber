package config

import (
	"PZN/entities"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var DatabaseUri = "root:@tcp(127.0.0.1:3306)/go_fiber?charset=utf8mb4&parseTime=True&loc=Local"

func Connect() error {

	var err error

	Database, err = gorm.Open(mysql.Open(DatabaseUri), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic("Database tidak tehubung")
	}

	err = Database.AutoMigrate(&entities.User{})
	if err != nil {
		return err
	}

	return nil

}
