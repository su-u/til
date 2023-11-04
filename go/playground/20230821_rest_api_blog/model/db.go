package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Post struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author string `json:"author"`
}

func newDb() (db *gorm.DB) {
	db, err := gorm.Open("sqlite", "data/database.db")
	if err != nil {
		panic(err)
	}
	return db
}

func GetPosts(limit int) ([]Post, error) {
	db := newDb()

	var posts []Post
	db.Raw("SELECT id, title, body, author FROM posts LIMIT ?", limit).Scan(&posts)
	return posts, db.Error
}

func Retrieve(id int) (Post, error) {
	db := newDb()
	post := Post{}

	db.Raw("SELECT id, title, body, author FROM posts WHERE id = ?", id).Scan(&post)
	return post, db.Error
}

func (post *Post) Create() error {
	db := newDb()
	db.Raw("INSERT INTO posts (title, body, author) VALUES (?, ?, ?) RETURNING id", post.Title, post.Body, post.Author)
	return db.Error
}

func (post *Post) Update() error {
	db := newDb()
	db.Raw("UPDATE posts SET title = ?, body = ?, author = ? WHERE id = ?", post.Title, post.Body, post.Author, post.Id)
	return db.Error
}

func (post *Post) Delete() error {
	db := newDb()
	db.Raw("DELETE FROM posts WHERE id = ?", post.Id)
	return db.Error
}
