package factory

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	cd "github.com/ivands26/Learn-Gin/feature/comment/data"
	cDev "github.com/ivands26/Learn-Gin/feature/comment/delivery"
	cu "github.com/ivands26/Learn-Gin/feature/comment/usecase"
)

func InitFactory(r *gin.Engine, db *gorm.DB) {
	cData := cd.NewData(db)
	cUsecase := cu.NewUseCase(cData)
	cHandler := cDev.NewHandler(cUsecase)
	cDev.RouteComment(r, cHandler)

}
