package maps

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/WillMcCall/Ten2/db/countries"
)

type GeocodeResponse struct {
	Results []struct {
		Geometry struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
		} `json:"geometry"`
	} `json:"results"`
}

func getCountryCoordinates(countryName string) (float64, float64, error) {
	apiKey := os.Getenv("GOOGLE_MAPS_KEY")
	fmt.Println(apiKey)

	// Encode the country name for use in the URL
	query := url.QueryEscape(countryName)
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?address=%s&key=%s", query, apiKey)
	fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	// Check if the response status is OK
	if resp.StatusCode != http.StatusOK {
		return 0, 0, fmt.Errorf("HTTP request failed with status: %s", resp.Status)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, err
	}

	// Print the raw response body for debugging
	fmt.Printf("API Response Body: %s\n", body)

	// Decode the JSON response
	var result GeocodeResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, 0, err
	}

	if len(result.Results) > 0 {
		return result.Results[0].Geometry.Location.Lat, result.Results[0].Geometry.Location.Lng, nil
	}

	return 0, 0, fmt.Errorf("no results found for %s", countryName)
}

func FormatCountryMapData(countries []countries.Country, iso3Code string) MapData {
	var lat, lng float64
	var zoom float32
	var found bool

	iso3 := []string{}
	hoverTexts := []string{}
	is_country := []float32{}

	for _, country := range countries {
		iso3 = append(iso3, country.ISO3)

		if country.ISO3 == iso3Code && !found {
			is_country = append(is_country, 1)

			var err error
			lat, lng, err = getCountryCoordinates(country.Name)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				found = true // Stop updating lat/lng once we get a valid response
			}

			// Adjusts the zoom based on country's population
			if country.Population > 200000000 {
				zoom = 3
			} else if country.Population > 50000000 {
				zoom = 4
			} else if country.Population < 250000 {
				zoom = 6
			} else {
				zoom = 5
			}

			if country.Name == "Russia" {
				zoom = 2
			}
		} else {
			is_country = append(is_country, 0)
		}

		hoverText := fmt.Sprintf("<b>%s</b><br><br>%.2f%% Evangelical<br>%.2f%% Christian Adherent",
			country.Name, country.PercentEvangelical, country.PercentChristianity)
		hoverTexts = append(hoverTexts, hoverText)
	}

	if !found {
		lat, lng = 0, 0 // Fallback to default values if no country is found
		zoom = 0
	}

	// Custom colorscale for the map
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
	data.Layout.Geo.Projection.Scale = zoom

	// Convert lat/lng to float32 before assigning
	data.Layout.Geo.Center.Latitude = float32(lat)
	data.Layout.Geo.Center.Longitude = float32(lng)

	return data
}
