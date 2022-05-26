package database

import (
	"fmt"
	"gin+jwt+postgres/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	localhost := config("DB_HOST")
	dbName := config("DB_NAME")
	dbUser := config("DB_USER")
	dbPassword := config("DB_PASSWORD")
	fmt.Printf("pass:%v", dbPassword)
	dbPort := config("DB_PORT")
	// if err != nil {
	// 	panic("cant convert to number")
	// }

	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable ", localhost, dbUser, dbPassword, dbName, dbPort)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("cant  connect database")
	}

	// DB.AutoMigrate{model.User{}}
	DB.AutoMigrate(&model.User{})
}

func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		panic("error close database")
	}
	// Close
	sqlDB.Close()
}
