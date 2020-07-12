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
	apiKey       = "3b728b75-36ab-4201-a229-bbc3266326dc"
)

type CatsData struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Temperament string `json:"temperament"`
	Origin      string `json:"origin"`
	Description string `json:"description"`
}

type CatsImages struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}

type ImagesCategories struct {
	Id   int8   `json:"id"`
	Name string `json:"name"`
}

func main() {
	FillDBWithCatsInfo()
	FillDBWithCatsImages()
	FillDBWithCatsImagesWithHats()
	FillDBWithCatsImagesWithGlasses()
}

// ===== REQUESTS CODES =====

func FillDBWithCatsInfo() {
	fmt.Println("Fetching cats details")
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
	fmt.Println("Filling DB with cats content")
	for _, catInfo := range myCatsData {
		InsertCatInfo(catInfo)
	}
}

func FillDBWithCatsImages() {
	fmt.Println("Fetching cats images")
	myCats := FetchCatsIDName()
	var myCatImages []CatsImages
	for _, cat := range myCats {
		client := &http.Client{}
		targetBreedRequest := fmt.Sprintf("https://api.thecatapi.com/v1/images/search?breed_id=%s&limit=3&mime_types=jpg,png", cat.Id)
		req, err := http.NewRequest("GET", targetBreedRequest, nil)
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
		errJSON := json.Unmarshal(bodyText, &myCatImages)
		if errJSON != nil {
			log.Println(errJSON)
		}
		fmt.Println(fmt.Sprintf("Filling DB with %s images", cat.Name))
		for _, catImageInfo := range myCatImages {
			InsertCatImage(catImageInfo, cat.Id, cat.Name)
		}
	}
}

func FillDBWithCatsImagesWithHats() {
	fmt.Println("Fetching images of cats with hats")
	var myCatImages []CatsImages
	client := &http.Client{}
	targetBreedRequest := fmt.Sprintf("https://api.thecatapi.com/v1/images/search?limit=3&mime_types=jpg,png&category_ids=%d", GetCatCategories("hats"))
	req, err := http.NewRequest("GET", targetBreedRequest, nil)
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
	errJSON := json.Unmarshal(bodyText, &myCatImages)
	if errJSON != nil {
		log.Println(errJSON)
	}
	fmt.Println("Filling DB with images of cats with hats")
	for _, catImageInfo := range myCatImages {
		InsertStylishCatImage(catImageInfo, false, true)
	}

}

func FillDBWithCatsImagesWithGlasses() {
	fmt.Println("Fetching images of cats with glasses")
	var myCatImages []CatsImages
	client := &http.Client{}
	targetBreedRequest := fmt.Sprintf("https://api.thecatapi.com/v1/images/search?limit=3&mime_types=jpg,png&category_ids=%d", GetCatCategories("sunglasses"))
	req, err := http.NewRequest("GET", targetBreedRequest, nil)
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
	errJSON := json.Unmarshal(bodyText, &myCatImages)
	if errJSON != nil {
		log.Println(errJSON)
	}
	fmt.Println("Filling DB with images of cats with glasses")
	for _, catImageInfo := range myCatImages {
		InsertStylishCatImage(catImageInfo, true, false)
	}
}

func GetCatCategories(targetCategory string) int8 {
	var imagesCategories []ImagesCategories
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.thecatapi.com/v1/categories", nil)
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
	errJSON := json.Unmarshal(bodyText, &imagesCategories)
	if errJSON != nil {
		log.Println(errJSON)
	}
	for _, categoriesInfo := range imagesCategories {
		if categoriesInfo.Name == targetCategory {
			return categoriesInfo.Id
		}
	}
	return -1
}

// ===== DB CODES =====

func CreateDBLink() *sql.DB {
	db, err := sql.Open("mysql", "root:z5dOucrrYHXvUNTHDqcz@tcp(mysql_db:3306)/cats_api")
	// db, err := sql.Open("mysql", "root:z5dOucrrYHXvUNTHDqcz@tcp(localhost:6603)/cats_api")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// ==INSERTS==

func InsertCatInfo(cat_info CatsData) {
	stmt, err := catsDatabase.Prepare("INSERT INTO cats_breeds(id, breed_name, temperament, origin, breed_description) VALUES (?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.
	_, err = stmt.Exec(cat_info.Id, cat_info.Name, cat_info.Temperament, cat_info.Origin, cat_info.Description)
	if err != nil {
		log.Fatal(err)
	}

}

func InsertCatImage(cat_image CatsImages, breed_id string, breed_name string) {
	stmt, err := catsDatabase.Prepare("INSERT INTO cats_images(id, breed_id, breed_name, image_url) VALUES (?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.
	_, err = stmt.Exec(cat_image.Id, breed_id, breed_name, cat_image.Url)
	if err != nil {
		log.Fatal(err)
	}
}

func InsertStylishCatImage(cat_image CatsImages, has_glasses bool, has_hat bool) {
	stmt, err := catsDatabase.Prepare("INSERT INTO stylish_cats_images(id, image_url, has_glasses, has_hat) VALUES (?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.
	_, err = stmt.Exec(cat_image.Id, cat_image.Url, has_glasses, has_hat)
	if err != nil {
		log.Fatal(err)
	}
}

// ==SELECT==
func FetchCatsIDName() []CatsData {
	cats := make([]CatsData, 0)
	rows, err := catsDatabase.Query("SELECT ID, breed_name FROM cats_breeds")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var catInfo CatsData

		err = rows.Scan(&catInfo.Id, &catInfo.Name)
		if err != nil {
			log.Fatal(err)
		}
		cats = append(cats, catInfo)
	}
	return cats
}
