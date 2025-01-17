package main

import (
	"html/template"
	"net/http"

	"github.com/WillMcCall/Ten2/helpers/maps"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Add `safeJS` as a template function
	router.SetFuncMap(template.FuncMap{
		"safeJS": func(s string) template.JS {
			return template.JS(s)
		},
	})

	router.LoadHTMLGlob("/var/www/will-mccall.com/ten2/templates/*")

	// Loads styles
	router.Static("/styles", "./styles")

	router.GET("/", func(c *gin.Context) {
		// Generate map JSON data
		mapJSON := getMapJSON()

		// Render the template with the map JSON
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"mapJSON": mapJSON,
		})
	})

	router.GET("/countries/ROU", func(c *gin.Context) {
		mapJSON := getCountryJSON("ROU")

		c.HTML(http.StatusOK, "rou.tmpl", gin.H{
			"mapJSON": mapJSON,
		})
	})

	router.Run(":8080")
}

func getMapJSON() string {
	countries := maps.GetAllCountries()
	mapData := maps.FormatMapData(countries)
	mapJSON := maps.ConvertMapToJSON(mapData)

	return string(mapJSON)
}

func getCountryJSON(iso3Code string) string {
	countries := maps.GetAllCountries()
	mapData := maps.FormatCountryMapData(countries, iso3Code)
	mapJSON := maps.ConvertMapToJSON(mapData)

	return string(mapJSON)
}
