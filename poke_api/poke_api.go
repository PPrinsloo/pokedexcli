package poke_api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokeApiResponse struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous any        `json:"previous"`
	Results  []Location `json:"results"`
}

var offset = 0
var limit = 10

func Locations() {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location?limit=%d&offset=%d", limit, offset)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	var response PokeApiResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total locations: %d\n", response.Count)
	for _, location := range response.Results {
		fmt.Println(location.Name)
	}
	offset += limit
}

func LocationsBack() {
	if offset-limit < 0 {
		fmt.Println("You are at the beginning of the list.")
		return
	}
	offset -= limit
	Locations()
}
