package helpers

import (
	"github.com/WillMcCall/Ten2/db"
	"github.com/WillMcCall/Ten2/db/countries"
)

func GetCountry(iso3 string) countries.Country {
	db := db.OpenConnection()
	country := *countries.GetByISO3(db, iso3)

	return country
}
