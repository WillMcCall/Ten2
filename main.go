package main

import (
	"html/template"
	"net/http"

	"github.com/WillMcCall/Ten2/helpers"
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

	router.LoadHTMLGlob("templates/*")
	router.Static("/styles", "./styles")

	router.GET("/", func(c *gin.Context) {
		// Generate map JSON data
		mapJSON := string(getMapJSON())

		// Render the template with the map JSON
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"mapJSON": mapJSON,
		})
	})

	router.Run(":8080")
}

func getMapJSON() []byte {
	countries := helpers.GetAllCountries()
	mapData := helpers.FormatMapData(countries)
	mapJSON := helpers.ConvertMapToJSON(mapData)

	return mapJSON
}
