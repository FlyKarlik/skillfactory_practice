package model

var (
	increment int = 0
)

type Authors struct {
	Id   int    `json:"-" bson:"id"`
	Name string `json:"name" bson:"name"`
}

type Posts struct {
	Id         int    `json:"id" bson:"id"`
	Author_id  int    `json:"author_id" bson:"author_id"`
	AuthorName string `json:"author_name" bson:"author_name"`
	Title      string `json:"title" bson:"title"`
	Content    string `json:"content" bson:"content"`
	Created    int    `json:"created" bson:"created"`
}

type UpdatePosts struct {
	UserId     int
	TaskId     int    `json:"task_id" bson:"id"`
	Title      string `json:"title" bson:"title"`
	Content    string `json:"content" bson:"content"`
	TypeUpdate string
}

func Next() int {
	increment++
	return increment
}
