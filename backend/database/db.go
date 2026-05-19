package database

import (
	"log"
	"os"

	"github.com/P-Raphat/2026-AMASS-Quiz-ThaiBev/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "devlocal:dev@2026@tcp(localhost:3307)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("failed to migrate:", err)
	}

	DB = db
}
