package repositories_test

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"goapps/models"
	"goapps/repositories"
)

var testCaseComments1 = models.Comment{
	CommentID: 1,
	ArticleID: 1,
	Message:   "test case comments 1",
}
var testCaseComments2 = models.Comment{
	CommentID: 2,
	ArticleID: 1,
	Message:   "test case comments 2",
}
var testCaseComments3 = models.Comment{
	CommentID: 3,
	ArticleID: 1,
	Message:   "test case comments 3",
}
var testCaseComments4 = models.Comment{
	CommentID: 4,
	ArticleID: 2,
	Message:   "test case comments 4",
}

func createCommentsObject() repositories.CommentInterface {
	db, err := sql.Open("sqlite3", testDBName)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := db.Conn(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return repositories.NewComments(conn)
}

func Test_CommentsInsert(t *testing.T) {
	obj := createCommentsObject()

	testCaseCommentsList := []models.Comment{
		testCaseComments1, testCaseComments2, testCaseComments3, testCaseComments4,
	}

	for _, testCase := range testCaseCommentsList {
		if err := obj.Insert(testCase); err != nil {
			t.Fatal(err)
		}
	}
}

func Test_FindByArticleKey(t *testing.T) {
	obj := createCommentsObject()

	rets, err := obj.FindByArticleId(testCase1.ID)
	if err != nil {
		t.Fatal(err)
	}

	if len(rets) != 3 {
		t.Errorf("Failed find by article key, %v", rets)
	}

	rets, err = obj.FindByArticleId(testCase2.ID)
	if err != nil {
		t.Fatal(err)
	}

	if len(rets) != 1 {
		t.Errorf("Failed find by article key, %v", rets)
	}
}
