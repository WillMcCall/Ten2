package countries

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/WillMcCall/Ten2/db"
)

func Init() {
	db := db.OpenConnection()
	defer db.Close()

	err := createTable(db)
	if err != nil {
		log.Fatal(err)
	}

	jsonBody := grabJSON()
	countries := convertJSON(jsonBody)

	for _, country := range countries {
		insert(db, country)
	}
}

func grabJSON() []byte {
	apiKey := os.Getenv("JOSHUA_PROJECT_KEY")
	url := "https://api.joshuaproject.net/v1/countries.json?api_key" + apiKey + "=b1dea565d8d0%20&limit=500&page=1"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("failed create request to Joshua Project API: %v", err)
	}

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("request to Joshua Project API failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("request failed with status: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("failed to read response body: %v", err)
	}

	fmt.Println("Successfully accessed Joshua Project API")
	return body
}

func convertJSON(jsonData []byte) []Country {
	var countries []Country

	err := json.Unmarshal(jsonData, &countries)
	if err != nil {
		log.Fatalf("error converting json: %v", err)
	}
	return countries
}
