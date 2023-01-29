package services_test

import (
	"database/sql"
	"fmt"
	"testing"

	"goapps/models"
	"goapps/services"
)

type CommentInterfaceMock struct{}

func (m *CommentInterfaceMock) Insert(comment models.Comment) error {
	return nil
}

func (m *CommentInterfaceMock) FindByArticleId(articleId int) ([]models.Comment, error) {
	return nil, nil
}

var insertErrMsg string = "mock, insert error."

type CommentInterfaceErrMock struct{}

func (m *CommentInterfaceErrMock) Insert(comment models.Comment) error {
	return fmt.Errorf(insertErrMsg)
}

func (m *CommentInterfaceErrMock) FindByArticleId(articleId int) ([]models.Comment, error) {
	return nil, sql.ErrNoRows
}

type CommentInterfaceErrMockError struct{}

func (m *CommentInterfaceErrMockError) Insert(comment models.Comment) error {
	return fmt.Errorf(insertErrMsg)
}

func (m *CommentInterfaceErrMockError) FindByArticleId(articleId int) ([]models.Comment, error) {
	return nil, fmt.Errorf("comments table select error message.")
}

func Test_InsertComment(t *testing.T) {
	commentService := &services.CommentService{
		CommentsInterface: &CommentInterfaceMock{},
	}

	err := commentService.InsertComment(models.Comment{})
	if err != nil {
		t.Fatal(err)
	}
}

func Test_InsertCommentErr(t *testing.T) {
	commentService := &services.CommentService{
		CommentsInterface: &CommentInterfaceErrMock{},
	}

	err := commentService.InsertComment(models.Comment{})
	if err == nil {
		t.Fatal(err)
	}
	if err.Error() != insertErrMsg {
		t.Fatal(err)
	}
}
