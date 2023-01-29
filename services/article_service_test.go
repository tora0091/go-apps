package services_test

import (
	"database/sql"
	"fmt"
	"goapps/models"
	"goapps/services"
	"testing"
	"time"
)

type ArticlesInterfaceMock struct{}

func (*ArticlesInterfaceMock) FindAll() ([]models.Article, error) {
	return nil, nil
}
func (*ArticlesInterfaceMock) FindById(articleId int) (*models.Article, error) {
	return nil, nil
}
func (*ArticlesInterfaceMock) Insert(article models.Article) error {
	return nil
}
func (*ArticlesInterfaceMock) UpdateNiceNum(articleId int) error {
	return nil
}

var articlesInsertErrMsg string = "Error, Mock Insert error message."
var articlesUpdateNiceNum string = "Error, Mock Update error message."

var testArticle models.Article = models.Article{
	ID:       1,
	Title:    "title",
	Contents: "contents",
	UserName: "username",
	NiceNum:  1,
	Comments: nil,
	CreateAt: time.Now(),
}

var testArticle2 models.Article = models.Article{
	ID:       2,
	Title:    "title2",
	Contents: "contents2",
	UserName: "username2",
	NiceNum:  2,
	Comments: nil,
	CreateAt: time.Now(),
}

var testArticle3 models.Article = models.Article{
	ID:       3,
	Title:    "title3",
	Contents: "contents3",
	UserName: "username3",
	NiceNum:  3,
	Comments: nil,
	CreateAt: time.Now(),
}

type ArticlesInterfaceErrMock struct{}

func (*ArticlesInterfaceErrMock) FindAll() ([]models.Article, error) {
	var testArticleList []models.Article = []models.Article{
		testArticle, testArticle2, testArticle3,
	}
	return testArticleList, nil
}
func (*ArticlesInterfaceErrMock) FindById(articleId int) (*models.Article, error) {
	return &testArticle, nil
}
func (*ArticlesInterfaceErrMock) Insert(article models.Article) error {
	return fmt.Errorf(articlesInsertErrMsg)
}
func (*ArticlesInterfaceErrMock) UpdateNiceNum(articleId int) error {
	return fmt.Errorf(articlesUpdateNiceNum)
}

type ArticlesInterfaceErrMockError struct{}

func (*ArticlesInterfaceErrMockError) FindAll() ([]models.Article, error) {
	return nil, fmt.Errorf("test error message.")
}
func (*ArticlesInterfaceErrMockError) FindById(articleId int) (*models.Article, error) {
	return nil, fmt.Errorf("test error message.")
}
func (*ArticlesInterfaceErrMockError) Insert(article models.Article) error {
	return nil
}
func (*ArticlesInterfaceErrMockError) UpdateNiceNum(articleId int) error {
	return nil
}

type ArticlesInterfaceErrMockNoRows struct{}

func (*ArticlesInterfaceErrMockNoRows) FindAll() ([]models.Article, error) {
	return nil, sql.ErrNoRows
}
func (*ArticlesInterfaceErrMockNoRows) FindById(articleId int) (*models.Article, error) {
	return nil, sql.ErrNoRows
}
func (*ArticlesInterfaceErrMockNoRows) Insert(article models.Article) error {
	return nil
}
func (*ArticlesInterfaceErrMockNoRows) UpdateNiceNum(articleId int) error {
	return nil
}

func Test_GetArticleDetailOK(t *testing.T) {
	targetService := services.ArticleService{
		ArticlesInterface: &ArticlesInterfaceErrMock{},
		CommentsInterface: &CommentInterfaceErrMock{},
	}

	article, err := targetService.GetArticleDetail(1)
	if err != nil {
		t.Fatal(err)
	}

	if article.ID != testArticle.ID {
		t.Fatal("not equal, id")
	}
	if article.UserName != testArticle.UserName {
		t.Fatal("not equal, user name")
	}
}

func Test_GetArticleDetailNoRows(t *testing.T) {
	targetService := services.ArticleService{
		ArticlesInterface: &ArticlesInterfaceErrMockNoRows{},
		CommentsInterface: nil,
	}

	_, err := targetService.GetArticleDetail(1)
	if err == nil {
		t.Fatal("no error")
	}
	if err.Error() != "sql: no rows in result set" {
		t.Fatal(err)
	}
}

func Test_GetArticleDetailError(t *testing.T) {
	targetService := services.ArticleService{
		ArticlesInterface: &ArticlesInterfaceErrMockError{},
		CommentsInterface: nil,
	}

	_, err := targetService.GetArticleDetail(1)
	if err == nil {
		t.Fatal("no error")
	}
	if err.Error() != "test error message." {
		t.Fatal(err)
	}
}

func Test_GetArticleDetailCommentsNoRows(t *testing.T) {
	targetService := services.ArticleService{
		ArticlesInterface: &ArticlesInterfaceErrMock{},
		CommentsInterface: &CommentInterfaceErrMock{},
	}

	article, err := targetService.GetArticleDetail(1)
	if err != nil {
		t.Fatal(err)
	}

	if article.Comments != nil {
		t.Fatal("article comments not nil")
	}
}

func Test_GetArticleDetailCommentsError(t *testing.T) {
	targetService := services.ArticleService{
		ArticlesInterface: &ArticlesInterfaceErrMock{},
		CommentsInterface: &CommentInterfaceErrMockError{},
	}

	_, err := targetService.GetArticleDetail(1)
	if err == nil {
		t.Fatal("no error message.")
	}

	if err.Error() != "comments table select error message." {
		t.Fatal(err)
	}
}

func Test_GetArticleList(t *testing.T) {
	targetService := services.ArticleService{
		ArticlesInterface: &ArticlesInterfaceErrMock{},
		CommentsInterface: nil,
	}

	articleList, err := targetService.GetArticleList()
	if err != nil {
		t.Fatal(err)
	}

	if len(articleList) != 3 {
		t.Fatal("failed get article list.")
	}
}

func Test_GetArticleListNoRows(t *testing.T) {
	targetService := services.ArticleService{
		ArticlesInterface: &ArticlesInterfaceErrMockNoRows{},
		CommentsInterface: nil,
	}

	_, err := targetService.GetArticleList()
	if err == nil {
		t.Fatal("no error.")
	}

	if err.Error() != "sql: no rows in result set" {
		t.Fatal(err)
	}
}

func Test_GetArticleListError(t *testing.T) {
	targetService := services.ArticleService{
		ArticlesInterface: &ArticlesInterfaceErrMockError{},
		CommentsInterface: nil,
	}

	_, err := targetService.GetArticleList()
	if err == nil {
		t.Fatal("no error.")
	}

	if err.Error() != "test error message." {
		t.Fatal(err)
	}
}

func Test_InsertArticle(t *testing.T) {
	targetService := services.ArticleService{
		ArticlesInterface: &ArticlesInterfaceMock{},
		CommentsInterface: nil,
	}
	if err := targetService.InsertArticle(models.Article{}); err != nil {
		t.Fatal(err)
	}

	targetService.ArticlesInterface = &ArticlesInterfaceErrMock{}
	err := targetService.InsertArticle(models.Article{})
	if err.Error() != articlesInsertErrMsg {
		t.Fatal(err)
	}
}

func Test_UpdateArticleNiceNum(t *testing.T) {
	targetService := services.ArticleService{
		ArticlesInterface: &ArticlesInterfaceMock{},
		CommentsInterface: nil,
	}
	if err := targetService.UpdateArticleNiceNum(1); err != nil {
		t.Fatal(err)
	}

	targetService.ArticlesInterface = &ArticlesInterfaceErrMock{}
	err := targetService.UpdateArticleNiceNum(1)
	if err.Error() != articlesUpdateNiceNum {
		t.Fatal(err)
	}
}
