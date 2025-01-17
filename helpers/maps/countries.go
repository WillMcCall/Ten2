package maps

import (
	"fmt"

	"github.com/WillMcCall/Ten2/db/countries"
)

func FormatCountryMapData(countries []countries.Country, iso3Code string) MapData {
	iso3 := []string{}
	hoverTexts := []string{}
	is_country := []float32{}

	for _, country := range countries {
		iso3 = append(iso3, country.ISO3)
		if country.ISO3 == iso3Code {
			is_country = append(is_country, 1)
		} else {
			is_country = append(is_country, 0)
		}

		hoverText := fmt.Sprintf("<b>%s<b><br><br>%.2f%% Evangelical<br>%.2f%% Christian", country.Name, country.PercentEvangelical, country.PercentChristianity)
		hoverTexts = append(hoverTexts, hoverText)
	}

	// Custom colorscale defined for the map
	customColorscale := [][]interface{}{
		{0.0, "rgb(136, 136, 136)"},
		{1.0, "rgb(45, 185, 255)"},
	}

	var data MapData
	data.Trace = append(data.Trace, Trace{})
	data.Trace[0].Type = "choropleth"
	data.Trace[0].Locations = iso3
	data.Trace[0].Values = is_country
	data.Trace[0].HoverText = hoverTexts
	data.Trace[0].HoverInfo = "text"
	data.Trace[0].Colorscale = customColorscale
	data.Trace[0].Showscale = false

	data.Layout.FigBgColor = "#eee"
	data.Layout.PlotBgColor = "#eee"
	data.Layout.Margin.Bottom = 5
	data.Layout.Margin.Top = 5
	data.Layout.Margin.Left = 0
	data.Layout.Margin.Right = 0

	data.Layout.Geo.ShowFrame = false
	data.Layout.Geo.ShowCoastlines = false
	data.Layout.Geo.Scope = "world"
	data.Layout.Geo.Projection.Type = "equirectangular"
	data.Layout.Geo.Projection.Scale = 3.9
	data.Layout.Geo.Center.Latitude = 50
	data.Layout.Geo.Center.Longitude = 22

	return data
}
