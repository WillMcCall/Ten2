package countries

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Country struct {
	Id                     string // 2 Letter FIPS 10-4 Country Code
	Name                   string
	ISO2                   string
	ISO3                   string
	Capital                string
	Population             int
	PercentBuddhism        float32
	PercentChristianity    float32
	PercentEthnicReligions float32
	PercentEvangelical     float32
	PercentHinduism        float32
	PercentIslam           float32
	PercentNonReligious    float32
	PercentOtherSmall      float32
	PercentUnknown         float32
}

func main() {
	db, err := sql.Open("sqlite3", "ten2.db")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to database")
	}

	createTable(db)

	defer db.Close()
}

func createTable(db *sql.DB) {
	_, err := db.Exec(`
		create table if not exists countries (
			id text primary key,
			name text not null,
			iso2 text unique not null,
			iso3 text unique not null,
			capital text,
			population integer,
			percent_buddhism real,
			percent_christianity real,
			percent_ethnicreligions real,
			percent_evangelical real,
			percent_hinduism real,
			percent_islam real,
			percent_nonreligious real,
			percent_othersmall real,
			percent_unknown real,
			been_there boolean CHECK(been_there = 0 OR been_there = 1),
			am_there boolean CHECK(am_there = 0 OR am_there = 1)
		);
	`) // Note to self: sqlite internally stores booleans as integers (0 or 1)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Countries table successfully created!")
	}
	createIndex(db, "iso2")
	createIndex(db, "iso3")
}

func createIndex(db *sql.DB, columnName string) {
	query := fmt.Sprintf("CREATE INDEX countries_%s_idx ON countries(%s)", columnName, columnName)

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Failed to create index on column %s: %v", columnName, err)
	} else {
		fmt.Printf("%s index successfully created!\n", columnName)
	}
}
