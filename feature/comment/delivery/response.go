package delivery

import "github.com/ivands26/Learn-Gin/domain"

type ResponseComment struct {
	CommentField string `json:"comment_field" form:"comment_field"`
	User         string `json:"user" form:"user"`
}

func FromModel(data []domain.Comment) []ResponseComment {
	res := []ResponseComment{}

	for _, val := range data {
		res = append(res, ResponseComment{
			CommentField: val.CommentField,
			User:         val.User,
		})
	}
	return res
}
