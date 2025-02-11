package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/WillMcCall/Ten2/helpers"
	"github.com/WillMcCall/Ten2/helpers/maps"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	router := gin.Default()

	router.SetFuncMap(template.FuncMap{
		"safeJS": func(s string) template.JS {
			return template.JS(s)
		},
		"formatFloat": func(n float32) string {
			return fmt.Sprintf("%.2f", n)
		},
		"formatNumberWithCommas": func(n int) string {
			s := strconv.Itoa(n) // Convert number to string
			if n < 1000 {
				return s
			}

			// Insert commas
			var result []string
			for len(s) > 3 {
				result = append([]string{s[len(s)-3:]}, result...)
				s = s[:len(s)-3]
			}
			result = append([]string{s}, result...)

			return strings.Join(result, ",")
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
