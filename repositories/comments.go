package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"goapps/models"
)

// use-case: CommentInterface
type CommentInterface interface {
	Insert(comment models.Comment) error
	FindByArticleId(articleId int) ([]models.Comment, error)
}

type Comments struct {
	conn *sql.Conn
}

func NewComments(conn *sql.Conn) CommentInterface {
	return &Comments{
		conn: conn,
	}
}

func (c *Comments) Insert(comment models.Comment) error {
	sqlText := fmt.Sprintf("INSERT INTO %s (comment_id, article_id, message, created_at) VALUES (?, ?, ?, ?);", COMMENT_TBL)

	stmt, err := c.conn.PrepareContext(context.Background(), sqlText)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(comment.CommentID, comment.ArticleID, comment.Message, time.Now().Format("2006-01-02T15:04:05Z07:00"))
	if err != nil {
		return err
	}
	return nil
}

func (c *Comments) FindByArticleId(articleId int) ([]models.Comment, error) {
	sqlText := fmt.Sprintf("SELECT comment_id, article_id, message, created_at FROM %s WHERE article_id = ?;", COMMENT_TBL)

	stmt, err := c.conn.PrepareContext(context.Background(), sqlText)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(articleId)
	if err != nil {
		return nil, err
	}

	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment

		err := rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.Message, &comment.CreateAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	if len(comments) == 0 {
		return nil, sql.ErrNoRows
	}
	return comments, nil
}
