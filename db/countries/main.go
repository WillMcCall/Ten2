package countries

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Country struct {
	Id                     int     `json:"id"`
	Name                   string  `json:"Ctry"`
	ISO2                   string  `json:"ISO2"`
	ISO3                   string  `json:"ISO3"`
	Capital                string  `json:"Capital"`
	Population             int     `json:"Population"`
	PercentBuddhism        float32 `json:"PercentBuddhism"`
	PercentChristianity    float32 `json:"PercentChristianity"`
	PercentEthnicReligions float32 `json:"PercentEthnicReligions"`
	PercentEvangelical     float32 `json:"PercentEvangelical"`
	PercentHinduism        float32 `json:"PercentHinduism"`
	PercentIslam           float32 `json:"PercentIslam"`
	PercentNonReligious    float32 `json:"PercentNonReligious"`
	PercentOtherSmall      float32 `json:"PercentOtherSmall"`
	PercentUnknown         float32 `json:"PercentUnknown"`
	BeenThere              bool    `json:"BeenThere"`
	AmThere                bool    `json:"AmThere"`
	GoingThere             bool    `json:"GoingThere"`
}

func createTable(db *sql.DB) error {
	_, err := db.Exec(`
		create table if not exists countries (
			id integer primary key autoincrement,
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
		fmt.Printf("error creating countries table: %v", err)
		return err
	} else {
		fmt.Println("Countries table successfully created!")
	}
	createIndex(db, "iso2")
	createIndex(db, "iso3")
	return nil
}

func createIndex(db *sql.DB, columnName string) error {
	query := fmt.Sprintf("CREATE INDEX IF NOT EXISTS countries_%s_idx ON countries(%s)", columnName, columnName)

	_, err := db.Exec(query)
	if err != nil {
		fmt.Errorf("failed to create index on column %s: %v", columnName, err)
		return err
	} else {
		fmt.Printf("%s index successfully created!\n", columnName)
		return nil
	}
}

func insert(db *sql.DB, country Country) error {
	query := `
		insert into countries (name, iso2, iso3, capital, population, percent_buddhism, percent_christianity, percent_ethnicreligions, 
		percent_evangelical, percent_hinduism, percent_islam, percent_nonreligious, percent_othersmall, percent_unknown, been_there, am_there, going_there)
		values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`

	_, err := db.Exec(query, country.Name, country.ISO2, country.ISO3, country.Capital, country.Population, country.PercentBuddhism,
		country.PercentChristianity, country.PercentEthnicReligions, country.PercentEvangelical, country.PercentHinduism, country.PercentIslam,
		country.PercentNonReligious, country.PercentOtherSmall, country.PercentUnknown, country.BeenThere, country.AmThere, country.GoingThere)

	if err != nil {
		fmt.Errorf("failed to insert %s: %v", country.Name, err)
		return err
	}
	fmt.Printf("%s inserted successfully!\n")
	return nil
}

func GetAll(db *sql.DB) *[]Country {
	rows, err := db.Query(`
		select id, name, iso2, iso3, capital, population, percent_buddhism, percent_christianity, percent_ethnicreligions, percent_evangelical, 
		percent_hinduism, percent_islam, percent_nonreligious, percent_othersmall, percent_unknown, been_there, am_there, going_there
		from countries
	`)
	if err != nil {
		fmt.Errorf("error querying countries table: %v", err)
		return nil
	}
	defer rows.Close()

	var countries []Country

	for rows.Next() {
		var country Country
		err := rows.Scan(&country.Id, &country.Name, &country.ISO2, &country.ISO3, &country.Capital, &country.Population, &country.PercentBuddhism,
			&country.PercentChristianity, &country.PercentEthnicReligions, &country.PercentEvangelical, &country.PercentHinduism,
			&country.PercentIslam, &country.PercentNonReligious, &country.PercentOtherSmall, &country.PercentUnknown, &country.BeenThere,
			&country.AmThere, &country.GoingThere)
		if err != nil {
			fmt.Errorf("error scanning data: %v", err)
			return nil
		}
		countries = append(countries, country)
	}

	// Check for errors iterating over rows
	if err := rows.Err(); err != nil {
		fmt.Errorf("row iteration failed: %v", err)
		return nil
	}

	return &countries
}

func GetByISO3(db *sql.DB, iso3 string) *Country {
	var country Country

	err := db.QueryRow(`
		select id, name, iso2, iso3, capital, population, percent_buddhism, percent_christianity, percent_ethnicreligions, percent_evangelical, 
		percent_hinduism, percent_islam, percent_nonreligious, percent_othersmall, percent_unknown, been_there, am_there, going_there
		from countries
		where iso3 = ?
	`, iso3).Scan(&country.Id, &country.Name, &country.ISO2, &country.ISO3, &country.Capital, &country.Population, &country.PercentBuddhism,
		&country.PercentChristianity, &country.PercentEthnicReligions, &country.PercentEvangelical, &country.PercentHinduism,
		&country.PercentIslam, &country.PercentNonReligious, &country.PercentOtherSmall, &country.PercentUnknown, &country.BeenThere,
		&country.AmThere, &country.GoingThere)

	if err != nil {
		fmt.Errorf("error querying %s: %v", iso3, err)
		return nil
	}

	return &country
}

func deleteAll(db *sql.DB) int {
	tx, err := db.Begin()
	if err != nil {
		fmt.Errorf("error starting the delete transaction: %v", err)
		return 0
	}

	result, err := tx.Exec("delete from countries")
	if err != nil {
		tx.Rollback()
		fmt.Errorf("delete failed: %v", err)
		return 0
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Errorf("failed to get affected rows: %v", err)
		return 0
	}

	return int(rowsAffected)
}

func deleteByISO3(db *sql.DB, iso3 string) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %v", err)
	}

	result, err := tx.Exec("DELETE FROM countries WHERE iso3 = ?", iso3)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete %s: %v", iso3, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error getting affected rows: %v", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}

	fmt.Printf("Successfully deleted %s (rows affected: %d)\n", iso3, rowsAffected)
	return nil
}
