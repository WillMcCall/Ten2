package countries

import (
	"database/sql"
	"fmt"
	"log"
	// _ "github.com/mattn/go-sqlite3"
)

type Country struct {
	Id   string // 2 Letter FIPS 10-4 Country Code
	Name string
	ISO2 string
	ISO3 string
}

func main() {
	db, err := sql.Open("sqlite3", "ten2.db")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to database")
	}
	defer db.Close()
}

func CreateTable(db *sql.DB) {
	// ToDo come up with better sql for this
	// Add Joshua Project Metrics
	// Add booleans for if I've been there yet or not
	// Figure out if you can constrain the length of text (like varchar(2))
	_, err := db.Exec(`
		create table if not exists countries (
			id text primary key,
			name text not null,
			code text unique not null
		);
	`)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Countries table successfully created!")
	}
}
