package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

var (
	apiKey = "<my_api_key>"
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
	router := mux.NewRouter()
	router.HandleFunc("/breeds", GetCatsBreeds).Methods("GET")
	router.HandleFunc("/breeds/{id}", GetCatsByID).Methods("GET")
	router.HandleFunc("/temperament/{temperament}", GetCatsByTemperament).Methods("GET")
	router.HandleFunc("/origin/{origin_id}", GetCatsByOrigin).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetCatsBreeds(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching all cats breeds")
	var myCatsData []CatsData
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
	json.NewEncoder(w).Encode(myCatsData)
}

func GetCatsByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(fmt.Sprintf("Fetching %s cat details", params["id"]))
	var myCatsData []CatsData
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.thecatapi.com/v1/breeds/search?q=%s", params["id"]), nil)
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
	json.NewEncoder(w).Encode(myCatsData)
}

func GetCatsByTemperament(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(fmt.Sprintf("Fetching cats breeds by %s temperament", params["temperament"]))
	var myCatsData []CatsData
	var returnData []CatsData
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
	for _, catInfo := range myCatsData {
		if strings.Contains(strings.ToLower(catInfo.Temperament), strings.ToLower(params["temperament"])) {
			returnData = append(returnData, catInfo)
		}
	}
	if returnData == nil {
		json.NewEncoder(w).Encode(make([]string, 0))
	} else {
		json.NewEncoder(w).Encode(returnData)
	}

}

func GetCatsByOrigin(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(fmt.Sprintf("Fetching cats breeds by %s origin", params["origin_id"]))
	var myCatsData []CatsData
	var returnData []CatsData
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
	for _, catInfo := range myCatsData {
		if strings.Contains(strings.ToLower(catInfo.Origin), strings.ToLower(params["origin_id"])) {
			returnData = append(returnData, catInfo)
		}
	}
	if returnData == nil {
		json.NewEncoder(w).Encode(make([]string, 0))
	} else {
		json.NewEncoder(w).Encode(returnData)
	}
}
