package main

import (
	"html/template"
	"net/http"

	"github.com/WillMcCall/Ten2/helpers"
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
			"Countries": maps.GetAllCountries(),
		})
	})

	router.GET("/countries/:iso3", func(c *gin.Context) {
		iso3 := c.Param("iso3")
		mapJSON := getCountryJSON(iso3)
		country := helpers.GetCountry(iso3)

		c.HTML(http.StatusOK, "pages/country", gin.H{
			"mapJSON": mapJSON,
			"country": country,
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
