package main

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ivands26/Learn-Gin/config"
	"github.com/ivands26/Learn-Gin/factory"
	"github.com/ivands26/Learn-Gin/infra/database"
)

func main() {
	cfg := config.GetConfig()
	db := database.InitDB(cfg)
	database.MigrateDB(db)
	r := gin.New()
	dsn := fmt.Sprintf(":%d", config.SERVERPORT)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "HEAD", "PATCH", "OPTIONS", "GET", "PUT"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	factory.InitFactory(r, db)

	r.Run(dsn)

}
