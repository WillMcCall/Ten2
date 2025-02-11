package maps

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/WillMcCall/Ten2/db"
	"github.com/WillMcCall/Ten2/db/countries"
)

type Trace struct {
	Type       string          `json:"type"`
	Locations  []string        `json:"locations"`
	Values     []float32       `json:"z"`
	HoverText  []string        `json:"text"`
	HoverInfo  string          `json:"hoverinfo"`
	Colorscale [][]interface{} `json:"colorscale"`
	Showscale  bool            `json:"showscale"`
}

type MapData struct {
	Trace  []Trace `json:"data"`
	Layout struct {
		FigBgColor  string `json:"paper_bgcolor"`
		PlotBgColor string `json:"plot_bgcolor"`
		Margin      struct {
			Left   int `json:"l"`
			Right  int `json:"r"`
			Bottom int `json:"b"`
			Top    int `json:"t"`
		} `json:"margin"`
		Geo struct {
			ShowFrame      bool   `json:"showframe"`
			ShowCoastlines bool   `json:"showcoastlines"`
			Scope          string `json:"scope"`
			Projection     struct {
				Type  string  `json:"type"`
				Scale float32 `json:"scale"`
			} `json:"projection"`
			Center struct {
				Latitude  float32 `json:"lat"`
				Longitude float32 `json:"lon"`
			} `json:"center"`
		} `json:"geo"`
	} `json:"layout"`
}

func GetAllCountries() []countries.Country {
	db := db.OpenConnection()
	countries := *countries.GetAll(db)

	return countries
}

func FormatMapData(countries []countries.Country) MapData {
	iso3 := []string{}
	hoverTexts := []string{}
	percent_evangelical := []float32{}

	for _, country := range countries {
		iso3 = append(iso3, country.ISO3)
		percent_evangelical = append(percent_evangelical, country.PercentEvangelical)
		hoverText := fmt.Sprintf("<b>%s<b><br><br>%.2f%% Evangelical<br>%.2f%% Christian Adherent",
			country.Name, country.PercentEvangelical, country.PercentChristianity)
		hoverTexts = append(hoverTexts, hoverText)
	}

	customColorscale := [][]interface{}{
		{0.0, "rgb(158, 10, 10)"},
		{0.05, "rgb(201, 73, 47)"},
		{0.3, "rgb(200, 173, 0)"},
		{0.6, "rgb(98, 215, 15)"},
		{1.0, "rgb(5, 243, 9)"},
	}

	var data MapData
	data.Trace = append(data.Trace, Trace{})
	data.Trace[0].Type = "choropleth"
	data.Trace[0].Locations = iso3
	data.Trace[0].Values = percent_evangelical
	data.Trace[0].HoverText = hoverTexts
	data.Trace[0].HoverInfo = "text"
	data.Trace[0].Colorscale = customColorscale
	data.Trace[0].Showscale = false

	data.Layout.FigBgColor = "#fff"
	data.Layout.PlotBgColor = "#fff"
	data.Layout.Margin.Bottom = 5
	data.Layout.Margin.Top = 5
	data.Layout.Margin.Left = 0
	data.Layout.Margin.Right = 0

	data.Layout.Geo.ShowFrame = false
	data.Layout.Geo.ShowCoastlines = false
	data.Layout.Geo.Scope = "world"
	data.Layout.Geo.Projection.Type = "equirectangular"
	data.Layout.Geo.Projection.Scale = 1
	data.Layout.Geo.Center.Latitude = 0
	data.Layout.Geo.Center.Longitude = 0

	return data
}

func ConvertMapToJSON(mapData MapData) []byte {
	jsonData, err := json.Marshal(mapData)
	if err != nil {
		log.Fatalf("error marshaling json: %v", err)
	}

	return jsonData
}
