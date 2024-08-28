package blackhatbook

import (
	"encoding/json"
	"log"
	"net/http"
	"testing"
)

type Status struct {
	Films   string `json:film`
	People  string `json:people`
	Planets string `json:planets`
}

/*
	"planets": "https://swapi.dev/api/planets/",
	"species": "https://swapi.dev/api/species/",
	"starships": "https://swapi.dev/api/starships/",
	"vehicles": "https://swapi.dev/api/vehicles/"
*/

func TestParsing(t *testing.T) {
	const URL = "https://swapi.dev/api/"

	res, err := http.Get(URL)
	if err != nil {
		log.Fatalln(err)
	}

	var status Status
	if err := json.NewDecoder(res.Body).Decode(&status); err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()
	log.Printf("%s -> %s\n", status.Films, status.People)
}
