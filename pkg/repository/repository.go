package repository

import (
	"skillfactory_project/model"

	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostsRepositoryPostgres interface {
	CreatePost(userId int, posts model.Posts) (int, error)
	GetAllPosts(userId int) ([]model.Posts, error)
	UpdatePost(updatePosts model.UpdatePosts) error
	DeletePost(deletePost model.Posts) error
}

type PostsRepositoryMongo interface {
	CreatePost(userId int, posts model.Posts) (int, error)
	GetAllPosts(userId int) ([]model.Posts, error)
	UpdatePost(updatePosts model.UpdatePosts) error
	DeletePost(deletePost model.Posts) error
}

type Repository struct {
	PostsRepositoryPostgres
	PostsRepositoryMongo
}

func NewRepository(db *sqlx.DB, mongo *mongo.Database) *Repository {
	return &Repository{
		PostsRepositoryPostgres: NewPostsPosgres(db),
		PostsRepositoryMongo:    NewPostsMongo(mongo),
	}
}
