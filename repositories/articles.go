package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"goapps/models"
)

// use-case: ArticlesInterface
type ArticlesInterface interface {
	FindAll() ([]models.Article, error)
	FindById(articleId int) (*models.Article, error)
	Insert(article models.Article) error
	UpdateNiceNum(articleId int) error
}

type Articles struct {
	conn *sql.Conn
}

func NewArticles(conn *sql.Conn) ArticlesInterface {
	return &Articles{
		conn: conn,
	}
}

func (a *Articles) FindAll() ([]models.Article, error) {
	sqlText := fmt.Sprintf("SELECT article_id, title, contents, username, nice, created_at FROM %s;", ARTICLE_TBL)

	rows, err := a.conn.QueryContext(context.Background(), sqlText)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []models.Article
	for rows.Next() {
		var article models.Article

		err := rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &article.CreateAt)
		if err != nil || err == sql.ErrNoRows {
			return nil, err
		}
		articles = append(articles, article)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	if len(articles) == 0 {
		return nil, sql.ErrNoRows
	}
	return articles, nil
}

func (a *Articles) FindById(articleId int) (*models.Article, error) {
	sqlText := fmt.Sprintf("SELECT article_id, title, contents, username, nice, created_at FROM %s WHERE article_id = ?;", ARTICLE_TBL)

	stmt, err := a.conn.PrepareContext(context.Background(), sqlText)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(articleId)

	var article models.Article
	err = row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &article.CreateAt)
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func (a *Articles) Insert(article models.Article) error {
	sqlText := fmt.Sprintf("INSERT INTO %s (article_id, title, contents, username, nice, created_at) VALUES (?, ?, ?, ?, ?, ?);", ARTICLE_TBL)

	stmt, err := a.conn.PrepareContext(context.Background(), sqlText)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(article.ID, article.Title, article.Contents, article.UserName, 0, time.Now().Format("2006-01-02T15:04:05Z07:00"))
	if err != nil {
		return err
	}
	return nil
}

func (a *Articles) UpdateNiceNum(articleId int) error {
	sqlText := fmt.Sprintf("UPDATE %s SET nice = nice + 1 WHERE article_id = ?;", ARTICLE_TBL)

	stmt, err := a.conn.PrepareContext(context.Background(), sqlText)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(articleId)
	if err != nil {
		return err
	}
	return nil
}
