package data

import "github.com/ivands26/Learn-Gin/domain"

type Comment struct {
	CommentField string
	User         string
}

func (c *Comment) ToDomain() domain.Comment {
	return domain.Comment{
		CommentField: c.CommentField,
		User:         c.User,
	}
}

func ParseToArr(arr []Comment) []domain.Comment {
	var res []domain.Comment
	for _, v := range arr {
		res = append(res, v.ToDomain())
	}
	return res
}

func FromDomain(dta domain.Comment) Comment {
	var res Comment
	res.CommentField = dta.CommentField
	res.User = dta.User
	return res
}
