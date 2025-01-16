package main

import (
	"fmt"

	"github.com/WillMcCall/Ten2/db"
	"github.com/WillMcCall/Ten2/db/countries"
)

func main() {
	// r := gin.Default()
	// r.GET("/test", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "test",
	// 	})
	// })
	// r.Run()

	db := db.OpenConnection()
	countries := countries.GetAll(db)

	fmt.Println(countries)
	fmt.Println(len(*countries))
	// I'VE DONE IT
}
