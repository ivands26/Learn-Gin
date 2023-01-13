package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/ivands26/Learn-Gin/domain"
)

func RouteComment(r *gin.Engine, ch domain.CommentHandler) {
	r.GET("/get", ch.GetComment())
}
