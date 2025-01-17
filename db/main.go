package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func OpenConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "ten2.db")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to database")
	return db
}
