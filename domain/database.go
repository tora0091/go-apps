package domain

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type ServeDatabase struct {
	DatabaseName string
}

func NewServeDatabase(databaseName string) *ServeDatabase {
	return &ServeDatabase{
		DatabaseName: databaseName,
	}
}

func (s *ServeDatabase) Run() {
	db, err := sql.Open("sqlite3", s.DatabaseName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createSql := `
	create table if not exists articles (
		article_id integer unsigned auto_increment primary key,
		title varchar(100) not null,
		contents text not null,
		username varchar(100) not null,
		nice integer not null,
		created_at datetime
	);
	create table if not exists comments (
		comment_id integer unsigned auto_increment primary key,
		article_id integer unsigned not null,
		message text not null,
		created_at datetime,
		foreign key (article_id) references articles(article_id)
	);`

	_, err = db.Exec(createSql)
	if err != nil {
		log.Fatal(err)
	}
}
