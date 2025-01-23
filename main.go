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

	router.Static("/static", "./static")  // Serves static files
	router.LoadHTMLGlob("templates/**/*") // Serves templates

	router.GET("/", func(c *gin.Context) {
		mapJSON := getMapJSON()

		c.HTML(http.StatusOK, "pages/home", gin.H{
			"mapJSON": mapJSON,
		})
	})

	router.GET("/countries", func(c *gin.Context) {

		c.HTML(http.StatusOK, "pages/countries", gin.H{
			"idk": "nothing",
		})
	})

	router.GET("/countries/ROU", func(c *gin.Context) {
		mapJSON := getCountryJSON("ROU")

		c.HTML(http.StatusOK, "pages/rou", gin.H{
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
