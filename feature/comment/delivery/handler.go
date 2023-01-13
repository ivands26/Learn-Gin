package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivands26/Learn-Gin/domain"
)

type commentHandler struct {
	chandler domain.CommentUC
}

func NewHandler(ch domain.CommentUC) domain.CommentHandler {
	return &commentHandler{
		chandler: ch,
	}
}

func (ch *commentHandler) GetComment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result, err := ch.chandler.GetCommentUC()

		if err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]interface{}{
				"code": 400,
				"msg":  "failed to get comment",
			})
			return
		}
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": 200,
			"msg":  "success get comment",
			"data": FromModel(result),
		})
	}

}
