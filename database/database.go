package database

import (
	"fmt"
	"os"
	"time"

	"github.com/welsonoktario/prakerja-unjuk-ketrampilan/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf(
		"%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, error := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if error != nil {
		panic("Error connecting to database: " + error.Error())
	}

	DB = db
	sqlDB, error := db.DB()

	if error != nil {
		panic(error)
	}

	db.AutoMigrate(models.User{}, models.Product{})
	sqlDB.SetConnMaxLifetime(time.Minute * 3)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(10)
	fmt.Println("Connected with Database")
}

func Get() *gorm.DB {
	return DB
}
