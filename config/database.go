package config

import (
	"fmt"
	"log"
	"os"

	"github.com/faizkhan-06/go-auth/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var DB *gorm.DB
func ConnectDb(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Dot Env file not found")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Faild to connect database")
	}

	DB = db

	fmt.Println("Database connected")

	db.AutoMigrate(&models.User{})
}