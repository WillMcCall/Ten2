package helpers

import (
	"encoding/json"
	"log"

	"github.com/WillMcCall/Ten2/db"
	"github.com/WillMcCall/Ten2/db/countries"
)

type MapData struct {
	Data  []Trace `json:"data"`
	Style Layout  `json:"layout"`
}

type Trace struct {
	Type       string    `json:"type"`
	Locations  []string  `json:"locations"`
	Values     []float32 `json:"z"`
	HoverText  []string  `json:"text"`
	Colorscale string    `json:"colorscale"`
}

type Layout struct {
	Title string `json:"title"`
}

func GetAllCountries() []countries.Country {
	db := db.OpenConnection()
	countries := *countries.GetAll(db)

	return countries
}

func FormatMapData(countries []countries.Country) MapData {
	iso3 := []string{}
	names := []string{}
	percent_evangelical := []float32{}

	for _, country := range countries {
		iso3 = append(iso3, country.ISO3)
		names = append(names, country.Name)
		percent_evangelical = append(percent_evangelical, country.PercentEvangelical)
	}

	var trace Trace
	var layout Layout
	var mapData MapData

	trace.Type = "choropleth"
	trace.Locations = iso3
	trace.Values = percent_evangelical
	trace.HoverText = names
	trace.Colorscale = "Electric"

	layout.Title = "Test Map"

	var traceSlice []Trace
	traceSlice = append(traceSlice, trace)

	mapData.Data = traceSlice
	mapData.Style = layout

	return mapData
}

func ConvertMapToJSON(mapData MapData) []byte {
	jsonData, err := json.Marshal(mapData)
	if err != nil {
		log.Fatalf("error marshaling json: %v", err)
	}

	return jsonData
}
