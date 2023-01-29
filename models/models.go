package models

import "time"

type Comment struct {
	CommentID int       `json:"comment_id"`
	ArticleID int       `json:"article_id"`
	Message   string    `json:"message"`
	CreateAt  time.Time `json:"create_at"`
}

type Article struct {
	ID       int       `json:"article_id"`
	Title    string    `json:"title"`
	Contents string    `json:"contents"`
	UserName string    `json:"user_name"`
	NiceNum  int       `json:"nice_num"`
	Comments []Comment `json:"comments"`
	CreateAt time.Time `json:"create_at"`
}
