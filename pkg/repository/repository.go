package repository

import (
	"skillfactory_project/model"
)

type PostsRepository interface {
	CreatePost(userId int, posts model.Posts) (int, error)
	GetAllPosts(userId int) ([]model.Posts, error)
	UpdatePost(updatePosts model.UpdatePosts) error
	DeletePost(deletePost model.Posts) error
}

type Repository struct {
	PostsRepository
}

func NewRepository() *Repository {
	return &Repository{}
}
