package services

import (
	"database/sql"

	"goapps/apperrors"
	"goapps/models"
	"goapps/repositories"
)

type ArticleService struct {
	ArticlesInterface repositories.ArticlesInterface
	CommentsInterface repositories.CommentInterface
}

func NewArticleService(conn *sql.Conn) *ArticleService {
	return &ArticleService{
		ArticlesInterface: repositories.NewArticles(conn),
		CommentsInterface: repositories.NewComments(conn),
	}
}

func (s *ArticleService) GetArticleDetail(articleId int) (*models.Article, error) {
	article, err := s.ArticlesInterface.FindById(articleId)
	if err == sql.ErrNoRows {
		return nil, apperrors.SelectNoRowsErr.Wrap("No Rows Articles table. func: GetArticleDetail", err)
	}
	if err != nil {
		return nil, apperrors.SelectErr.Wrap("Error Articles table select. func: GetArticleDetail", err)
	}

	comment, err := s.CommentsInterface.FindByArticleId(articleId)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, apperrors.SelectErr.Wrap("Error Comments table select. func: GetArticleDetail", err)
		}
	}

	article.Comments = comment

	return article, nil
}

func (s *ArticleService) GetArticleList() ([]models.Article, error) {
	articles, err := s.ArticlesInterface.FindAll()
	if err == sql.ErrNoRows {
		return nil, apperrors.SelectNoRowsErr.Wrap("No Rows Articles table. func: GetArticleList", err)
	}
	if err != nil {
		return nil, apperrors.SelectErr.Wrap("Error Articles table select. func: GetArticleDetail", err)
	}
	return articles, nil
}

func (s *ArticleService) InsertArticle(article models.Article) error {
	if err := s.ArticlesInterface.Insert(article); err != nil {
		return apperrors.InsertErr.Wrap("Error Articles table insert. func: InsertArticle", err)
	}
	return nil
}

func (s *ArticleService) UpdateArticleNiceNum(articleId int) error {
	if err := s.ArticlesInterface.UpdateNiceNum(articleId); err != nil {
		return apperrors.UpdateErr.Wrap("Error Articles table update. func: UpdateArticleNiceNum", err)
	}
	return nil
}
