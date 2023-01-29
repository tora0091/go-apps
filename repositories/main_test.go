package repositories_test

// go test ./...
// go test ./repositories
// go test -cover ./... -coverprofile=cover.out
// go tool cover -html=cover.out -o cover.html
import (
	"os"
	"testing"

	"goapps/domain"
)

const testDBName = "./article_test.db"

func setup() {
	domain.NewServeDatabase(testDBName).Run()
}

func teardown() {
	os.Remove(testDBName)
}

func TestMain(m *testing.M) {
	setup()

	m.Run()

	teardown()
}
