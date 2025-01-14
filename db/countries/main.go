package countries

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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
	BeenThere              bool
	AmThere                bool
	GoingThere             bool
}

func CreateTable(db *sql.DB) {
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
			am_there boolean CHECK(am_there = 0 OR am_there = 1),
			going_there boolean CHECK(going_there = 0 OR going_there = 1)
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
	query := fmt.Sprintf("CREATE INDEX IF NOT EXISTS countries_%s_idx ON countries(%s)", columnName, columnName)

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Failed to create index on column %s: %v", columnName, err)
	} else {
		fmt.Printf("%s index successfully created!\n", columnName)
	}
}

func insertCountry(db *sql.DB, country Country) {
	query := `
		insert into countries (id, name, iso2, iso3, capital, population, percent_buddhism, percent_christianity, percent_ethnicreligions, 
		percent_evangelical, percent_hinduism, percent_islam, percent_nonreligious, percent_othersmall, percent_unknown, been_there, am_there, going_there)
		values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`

	_, err := db.Exec(query, country.Id, country.Name, country.ISO2, country.ISO3, country.Capital, country.Population, country.PercentBuddhism,
		country.PercentChristianity, country.PercentEthnicReligions, country.PercentEvangelical, country.PercentHinduism, country.PercentIslam,
		country.PercentNonReligious, country.PercentOtherSmall, country.PercentUnknown, country.BeenThere, country.AmThere, country.GoingThere)

	if err != nil {
		log.Fatalf("Failed to insert %s: %v", country.Name, err)
	}
	log.Printf("%s inserted successfully!\n")
}

func grabCountriesJSON() []byte {
	apiKey := os.Getenv("JOSHUA_PROJECT_KEY")
	url := "https://api.joshuaproject.net/v1/countries.json?api_key" + apiKey + "=b1dea565d8d0%20&limit=500&page=1"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Failed create request to Joshua Project API: %v", err)
	}

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Request to Joshua Project API failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	fmt.Println("Successfully accessed Joshua Project API")
	return body
}
