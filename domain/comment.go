package domain

import "github.com/gin-gonic/gin"

type Comment struct {
	CommentField string
	User         string
}

type CommentData interface {
	GetCommentData() (data []Comment, err error)
}

type CommentUC interface {
	GetCommentUC() (data []Comment, err error)
}

type CommentHandler interface {
	GetComment() gin.HandlerFunc
}
