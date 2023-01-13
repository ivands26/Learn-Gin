package usecase

import (
	"fmt"

	"github.com/ivands26/Learn-Gin/domain"
)

type commentUC struct {
	commentData domain.CommentData
}

func NewUseCase(cd domain.CommentData) domain.CommentUC {
	return &commentUC{
		commentData: cd,
	}
}

func (cu *commentUC) GetCommentUC() (data []domain.Comment, err error) {
	result, err := cu.commentData.GetCommentData()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return result, nil
}
