package repositories_test

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"goapps/models"
	"goapps/repositories"
)

var testCase1 = models.Article{
	ID:       1,
	Title:    "Test Case 1 title",
	Contents: "Test Case 1 contents",
	UserName: "user name",
	NiceNum:  0,
}

var testCase2 = models.Article{
	ID:       2,
	Title:    "Test Case 2 title",
	Contents: "Test Case 2 contents",
	UserName: "user name",
	NiceNum:  0,
}

var testCase3 = models.Article{
	ID:       3,
	Title:    "Test Case 3 title",
	Contents: "Test Case 3 contents",
	UserName: "user name",
	NiceNum:  0,
}

var testCaseList = []models.Article{
	testCase1, testCase2, testCase3,
}

var noDataId = 99

func createObject() repositories.ArticlesInterface {
	db, err := sql.Open("sqlite3", testDBName)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := db.Conn(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return repositories.NewArticles(conn)
}

func Test_InsertArticle(t *testing.T) {
	a := createObject()

	for _, testCase := range testCaseList {
		ret := a.Insert(testCase)
		if ret != nil {
			t.Errorf("Failed insert article. %v", testCase)
		}
	}
}

func Test_InsertArticleDuplication(t *testing.T) {
	a := createObject()

	ret := a.Insert(testCase1)

	if ret == nil {
		t.Errorf("Failed insert article duplication. %v", testCase1)
	}
}

func Test_SelectArticleDetail(t *testing.T) {
	a := createObject()

	r, err := a.FindAll()
	if err != nil {
		t.Fatal(err)
	}

	if len(r) != len(testCaseList) {
		t.Errorf("Failed select article detail. %v", r)
	}
}

func Test_SelectArticleByKey(t *testing.T) {
	a := createObject()

	r, err := a.FindById(testCase3.ID)
	if err != nil {
		t.Fatal(err)
	}

	if r.ID != testCase3.ID {
		t.Errorf("Failed data ID, %v", r)
	}
	if r.Title != testCase3.Title {
		t.Errorf("Failed data Title, %v", r)
	}
	if r.Contents != testCase3.Contents {
		t.Errorf("Failed data Contents, %v", r)
	}
}

func Test_SelectArticleByKeyNoData(t *testing.T) {
	a := createObject()

	_, err := a.FindById(noDataId)
	if err == nil {
		t.Errorf("Failed no data, %v", err)
	}
}

func Test_UpdateNiceNum(t *testing.T) {
	a := createObject()

	if err := a.UpdateNiceNum(testCase2.ID); err != nil {
		t.Fatal(err)
	}

	r, err := a.FindById(testCase2.ID)
	if err != nil {
		t.Fatal(err)
	}

	if r.NiceNum != testCase2.NiceNum+1 {
		t.Errorf("Failed update nice num. %v", r)
	}
}
