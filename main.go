package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"goapps/domain"
	"goapps/route"
)

var databasePool *sql.DB

func init() {
	domain.NewServeDatabase("./article.db").Run()
}

func main() {
	conn, err := databaseConnection()
	if err != nil {
		log.Fatal(err)
	}

	r := route.NewCreateHandlers(conn).Routes()

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

	defer databasePool.Close()
}

func databaseConnection() (*sql.Conn, error) {
	if databasePool == nil {
		db, err := sql.Open("sqlite3", "./article.db")
		if err != nil {
			return nil, err
		}
		databasePool = db
	}
	ctx := context.Background()
	conn, err := databasePool.Conn(ctx)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
