package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var (
	catsDatabase = CreateDBLink()
)

type CatsWeightInner struct {
	Imperial string `json":"imperial"`
	Metric   string `json":"metric"`
}

type CatsData struct {
	Weight           CatsWeightInner `json:"weight"`
	Id               string          `json:"id"`
	Name             string          `json:"name"`
	CfaURL           string          `json:"cfa_url"`
	VetstreetURL     string          `json:"vetstreet_url"`
	VcahospitalsURL  string          `json:"vcahospitals_url"`
	Temperament      string          `json:"temperament"`
	Origin           string          `json:"origin"`
	CountryCodes     string          `json:"country_codes"`
	CountryCode      string          `json:"country_code"`
	Description      string          `json:"description"`
	LifeSpan         string          `json:"life_span"`
	Indoor           uint8           `json:"indoor"`
	Lap              uint8           `json:"lap"`
	AltNames         string          `json:"alt_names"`
	Adaptability     uint8           `json:"adaptability"`
	AffectionLevel   uint8           `json:"affection_level"`
	ChildFriendly    uint8           `json:"child_friendly"`
	DogFriendly      uint8           `json:"dog_friendly"`
	EnergyLevel      uint8           `json:"energy_level"`
	Grooming         uint8           `json:"grooming"`
	HealthIssues     uint8           `json:"health_issues"`
	Intelligence     uint8           `json:"intelligence"`
	SheddingLevel    uint8           `json:"shedding_level"`
	SocialNeeds      uint8           `json:"social_needs"`
	StrangerFriendly uint8           `json:"stranger_friendly"`
	Vocalisation     uint8           `json:"vocalisation"`
	Experimental     uint8           `json:"experimental"`
	Hairless         uint8           `json:"hairless"`
	Natural          uint8           `json:"natural"`
	Rare             uint8           `json:"rare"`
	Rex              uint8           `json:"rex"`
	SuppressedTail   uint8           `json:"suppressed_tail"`
	ShortLegs        uint8           `json:"short_legs"`
	WikipediaUrl     string          `json:"wikipedia_url"`
	Hypoallergenic   uint8           `json:"hypoallergenic"`
}

func main() {
	var myCatsData []CatsData
	apiKey := "<my_api_key>"
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.thecatapi.com/v1/breeds", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("api_key", apiKey)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	errJSON := json.Unmarshal(bodyText, &myCatsData)
	if errJSON != nil {
		log.Println(errJSON)
	}
	fmt.Println("Filling DB with JSON content")
	for _, catInfo := range myCatsData {
		RunInsert("INSERT INTO cats_breeds(id, breed_name, weight_imperial, weight_metric, cfa_url, vet_street_url, vca_hospitals_url, temperament, origin, country_codes, country_code, breed_description, lifes_span, indoor, lap, alt_name, adaptability, affection_level, child_friendly, dog_friendly, energy_level, grooming, health_issues, intelligence, shedding_level, social_needs, stranger_friendly, vocalisation, experimental, hairless, breed_natural, rare, rex, suppressed_tail, short_legs, wikipedia_url, hypoallergenic) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", catInfo)
	}
	RunSelect("SELECT ID, breed_name FROM cats_breeds")
}

func CreateDBLink() *sql.DB {
	db, err := sql.Open("mysql", "root:<my_password>@tcp(localhost:6603)/cats_api")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func RunInsert(query string, cat_info CatsData) {
	stmt, err := catsDatabase.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.
	_, err = stmt.Exec(cat_info.Id, cat_info.Name, cat_info.Weight.Imperial, cat_info.Weight.Metric, cat_info.CfaURL,
		cat_info.VetstreetURL, cat_info.VcahospitalsURL, cat_info.Temperament, cat_info.Origin, cat_info.CountryCodes,
		cat_info.CountryCode, cat_info.Description, cat_info.LifeSpan, cat_info.Indoor, cat_info.Lap, cat_info.AltNames,
		cat_info.Adaptability, cat_info.AffectionLevel, cat_info.ChildFriendly, cat_info.DogFriendly, cat_info.EnergyLevel,
		cat_info.Grooming, cat_info.HealthIssues, cat_info.Intelligence, cat_info.SheddingLevel, cat_info.SocialNeeds,
		cat_info.StrangerFriendly, cat_info.Vocalisation, cat_info.Experimental, cat_info.Hairless, cat_info.Natural, cat_info.Rare,
		cat_info.Rex, cat_info.SuppressedTail, cat_info.ShortLegs, cat_info.WikipediaUrl, cat_info.Hypoallergenic)
	if err != nil {
		log.Fatal(err)
	}

}

func RunSelect(query string) {
	rows, err := catsDatabase.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var cat_id string
		var cat_breed string

		err = rows.Scan(&cat_id, &cat_breed)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(cat_id, cat_breed)
	}
}
