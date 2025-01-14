package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sethgrid/pester"
)

type locationAreaEncounter struct {
	LocationArea struct {
		Name string
	} `json:"location_area"`
}

type pokemon struct {
	Name string `json:"name"`
	Locations []string `json:"locations"`
}

const pokeapiURL = "https://pokeapi.co/api/v2"

func main() {
	name := flag.String("name", "pikachu", "Pokemon name.")

	flag.Parse()

	httpClient := pester.New()

	httpResponse, err := httpClient.Get(fmt.Sprintf("%v/pokemon/%v/encounters", pokeapiURL, *name))
	if err != nil {
		fmt.Println("Error fetching pokemon from PokeAPI.")
		os.Exit(1)
	}

	bodyContent, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		fmt.Println("Error reading HTTP response body.")
		os.Exit(1)
	}

	var decodedContent []locationAreaEncounter
	err = json.Unmarshal(bodyContent, &decodedContent)
	if err != nil {
		fmt.Println("Error unmarshalling the JSON body content.")
		os.Exit(1)
	}

	poke := pokemon {
		Name: *name,
	}

	for _, location := range decodedContent {
		poke.Locations = append(poke.Locations, location.LocationArea.Name)
	}

	pokeJson, err := json.MarshalIndent(poke, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling the pokemon struct.")
		os.Exit(1)
	}

	fmt.Println(string(pokeJson))
}
