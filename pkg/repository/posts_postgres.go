package repository

import (
	"fmt"
	"skillfactory_project/model"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable = "authors"
	postsTable = "posts"
)

type PostsPostgres struct {
	db *sqlx.DB
}

func NewPostsPosgres(db *sqlx.DB) *PostsPostgres {
	return &PostsPostgres{db: db}
}

func (d *PostsPostgres) CreatePost(userId int, p model.Posts) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (title,content,author_id) VALUES ($1,$2,$3) RETURNING id", postsTable)
	rows := d.db.QueryRow(query, p.Title, p.Content, userId)
	if err := rows.Scan(&id); err != nil {
		return 0, err
	}
	return id, rows.Err()
}

func (d *PostsPostgres) GetAllPosts(userId int) ([]model.Posts, error) {
	var authorName string
	queryTask := fmt.Sprintf("SELECT id,author_id,title,content,created FROM %s WHERE author_id=$1", postsTable)
	queryUser := fmt.Sprintf("SELECT name FROM %s WHERE id=$1", usersTable)

	name := d.db.QueryRow(queryUser, userId)
	if err := name.Scan(&authorName); err != nil {
		return nil, err
	}

	rows, err := d.db.Query(queryTask, userId)
	if err != nil {
		return nil, err
	}

	var posts []model.Posts

	for rows.Next() {
		var p model.Posts
		err := rows.Scan(
			&p.Id,
			&p.Author_id,
			&p.Title,
			&p.Content,
			&p.Created,
		)
		if err != nil {
			return nil, err
		}

		p.AuthorName = authorName
		posts = append(posts, p)
	}

	return posts, rows.Err()
}

func (d *PostsPostgres) UpdatePost(updatePost model.UpdatePosts) error {
	if updatePost.TypeUpdate == "Both" {
		query := fmt.Sprintf("UPDATE %s SET title = $1, content = $2 WHERE id = $3 AND author_id = $4", postsTable)
		_, err := d.db.Exec(query, updatePost.Title, updatePost.Content, updatePost.TaskId, updatePost.UserId)
		if err != nil {
			return err
		}
	}

	if updatePost.TypeUpdate == "Title" {
		query := fmt.Sprintf("UPDATE %s SET title = $1 WHERE id = $2 AND author_id = $3", postsTable)
		_, err := d.db.Exec(query, updatePost.Title, updatePost.TaskId, updatePost.UserId)
		if err != nil {
			return err
		}
	}

	if updatePost.TypeUpdate == "Content" {
		query := fmt.Sprintf("UPDATE %s SET content = $1 WHERE id = $2 AND author_id = $3", postsTable)
		_, err := d.db.Exec(query, updatePost.Content, updatePost.TaskId, updatePost.UserId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *PostsPostgres) DeletePost(deletePost model.Posts) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 AND author_id=$2", postsTable)
	_, err := d.db.Exec(query, deletePost.Id, deletePost.Author_id)
	if err != nil {
		return err
	}
	return nil
}
