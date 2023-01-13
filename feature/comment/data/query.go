package data

import (
	"github.com/ivands26/Learn-Gin/domain"
	"gorm.io/gorm"
)

type commentData struct {
	db *gorm.DB
}

func NewData(db *gorm.DB) domain.CommentData {
	return &commentData{
		db: db,
	}
}

func (cd *commentData) GetCommentData() (data []domain.Comment, err error) {
	var comment []Comment
	result := cd.db.Table("comments").Find(&comment)

	if result.Error != nil {
		return nil, result.Error
	}

	return ParseToArr(comment), nil
}
