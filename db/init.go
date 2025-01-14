package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/WillMcCall/Ten2/countries"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "ten2.db")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to database")
	}
	defer db.Close()

	countries.CreateTable(db)
}
