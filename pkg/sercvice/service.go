package service

import (
	"skillfactory_project/model"
)

type PostsService interface {
	CreatePost(userId int, p model.Posts) (int, error)
	GetAllPosts(userId int) ([]model.Posts, error)
	UpdatePost(updatePost model.UpdatePosts) error
	DeletePost(deletePost model.Posts) error
}

type Sercive struct {
	PostsService
}

func NewService() *Sercive {
	return &Sercive{}
}
