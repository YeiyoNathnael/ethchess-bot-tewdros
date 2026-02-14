package db

import (
	"database/sql"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func Init(url string) (*sql.DB, error) {
	db, err := sql.Open("libsql", url)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
