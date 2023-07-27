package database

import (
	"fmt"
	"time"

	"github.com/welsonoktario/prakerja-unjuk-ketrampilan/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root@tcp(127.0.0.1:3306)/prakerja?charset=utf8mb4&parseTime=True&loc=Local"

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
