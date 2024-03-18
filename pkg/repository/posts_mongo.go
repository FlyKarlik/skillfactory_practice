package repository

import (
	"context"
	"skillfactory_project/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	authorsCollection = "users"
	postsCollection   = "posts"
)

type PostsMongo struct {
	db *mongo.Database
}

func NewPostsMongo(mongo *mongo.Database) *PostsMongo {
	return &PostsMongo{db: mongo}
}

func (d *PostsMongo) CreatePost(userId int, p model.Posts) (int, error) {
	collection := d.db.Collection(postsCollection)
	posts, err := d.FillPost(userId, p)
	if err != nil {
		return 0, err
	}
	_, err = collection.InsertOne(context.Background(), posts)
	if err != nil {
		return 0, nil
	}

	return posts.Id, nil
}

func (d *PostsMongo) GetAllPosts(userId int) ([]model.Posts, error) {
	collection := d.db.Collection(postsCollection)
	filter := bson.M{"author_id": userId}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	var posts []model.Posts

	for cur.Next(context.Background()) {
		var p model.Posts
		err := cur.Decode(&p)
		if err != nil {
			return nil, err
		}

		posts = append(posts, p)
	}

	return posts, cur.Err()
}

func (d *PostsMongo) UpdatePost(updatePost model.UpdatePosts) error {
	collection := d.db.Collection(postsCollection)
	filter := bson.M{"id": updatePost.TaskId}
	if updatePost.TypeUpdate == "Both" {
		_, err := collection.UpdateOne(context.Background(), filter, bson.M{"$set": bson.M{"title": updatePost.Title}, "content": updatePost.Content})
		if err != nil {
			return err
		}
	}

	if updatePost.TypeUpdate == "Title" {
		_, err := collection.UpdateOne(context.Background(), filter, bson.M{"$set": bson.M{"title": updatePost.Title}})
		if err != nil {
			return err
		}
	}

	if updatePost.TypeUpdate == "Content" {
		_, err := collection.UpdateOne(context.Background(), filter, bson.M{"$set": bson.M{"content": updatePost.Content}})
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *PostsMongo) DeletePost(deletePost model.Posts) error {
	collection := d.db.Collection(postsCollection)
	filter := bson.M{"id": deletePost.Id}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}

func (d *PostsMongo) FillPost(userId int, p model.Posts) (*model.Posts, error) {
	var user model.Authors
	userCollection := d.db.Collection(authorsCollection)
	if err := userCollection.FindOne(context.Background(), bson.M{"id": userId}).Decode(&user); err != nil {
		return nil, err
	}
	p.Id = model.Next()
	p.Author_id = user.Id
	p.AuthorName = user.Name
	p.Created = int(time.Now().Unix())
	return &p, nil
}
