package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ivands26/Learn-Gin/config"
	"github.com/ivands26/Learn-Gin/infra/database"
)

func main() {
	cfg := config.GetConfig()
	db := database.InitDB(cfg)
	database.MigrateDB(db)
	r := gin.New()
	dsn := fmt.Sprintf(":%d", config.SERVERPORT)

	r.Run(dsn)

}
