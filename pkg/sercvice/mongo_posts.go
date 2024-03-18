package service

import (
	"errors"
	"skillfactory_project/model"
	"skillfactory_project/pkg/repository"
)

type MongoPosts struct {
	repo repository.PostsRepositoryMongo
}

func NewMongoService(repo repository.PostsRepositoryMongo) *MongoPosts {
	return &MongoPosts{
		repo: repo,
	}
}

func (s *MongoPosts) CreatePost(userId int, p model.Posts) (int, error) {
	id, err := s.repo.CreatePost(userId, p)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *MongoPosts) GetAllPosts(userId int) ([]model.Posts, error) {
	posts, err := s.repo.GetAllPosts(userId)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *MongoPosts) UpdatePost(updatePost model.UpdatePosts) error {
	updates := s.Updates(updatePost)
	if updates.TypeUpdate == "No updates" {
		return errors.New("Nothing change")
	}
	if err := s.repo.UpdatePost(updates); err != nil {
		return err
	}

	return nil
}

func (s *MongoPosts) DeletePost(deletePost model.Posts) error {
	err := s.repo.DeletePost(deletePost)
	if err != nil {
		return err
	}

	return nil
}

func (s *MongoPosts) Updates(update model.UpdatePosts) model.UpdatePosts {
	if update.Title == "" && update.Content == "" {
		update.TypeUpdate = "No updates"
	}

	if update.Title != "" && update.Content != "" {
		update.TypeUpdate = "Both"
	}

	if update.Title != "" && update.Content == "" {
		update.TypeUpdate = "Title"
	}

	if update.Title == "" && update.Content != "" {
		update.TypeUpdate = "Content"
	}

	return update
}
