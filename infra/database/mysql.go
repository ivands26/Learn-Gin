package database

import (
	"fmt"
	"log"

	"github.com/ivands26/Learn-Gin/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	cData "github.com/ivands26/Learn-Gin/feature/comment/data"
)

func InitDB(cfg *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.Username,
		cfg.Password,
		cfg.Address,
		cfg.Port,
		cfg.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to Init Database")
		return nil
	}
	return db
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(cData.Comment{})
}
