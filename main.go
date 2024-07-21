package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var BASE_URL = "https://pokeapi.co/api/v2/"

type PokemonList struct {
	Count   int        `json:"count"`
	Next    string     `json:"next"`
	Prev    string     `json:"prev"`
	Results []Resource `json:"results"`
}

type Resource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
type Pokemon struct {
	Name   string        `json:"name"`
	ID     int           `json:"id"`
	Types  []PokemonType `json:"types"`
	Stats  []PokemonStat `json:"stats"`
	Weight int           `json:"weight"`
}
type PokemonType struct {
	Slot int      `json:"name"`
	Type Resource `json:"type"`
}
type PokemonStat struct {
	BaseStat int      `json:"base_stat"`
	Stat     Resource `json:"stat"`
}

func main() {
	fetchById("bulbasaur")

}

// Takes either an ID or name
func fetchById[T interface{ ~int | ~string }](searchTerm T) {
	res, err := http.Get(fmt.Sprintf("%s/pokemon/%v", BASE_URL, searchTerm))
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode == 404 {
		fmt.Println("Pokemon not found")
		return
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	var responseObj Pokemon
	json.Unmarshal(data, &responseObj)
	printData(responseObj)
}

func printData(pokemon Pokemon) {
	fmt.Println(pokemon.Name)
	fmt.Println(pokemon.ID)
	fmt.Println("Weight: ", pokemon.Weight)

	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Println(t.Type.Name)
	}
	fmt.Println("Stats:")
	for _, s := range pokemon.Stats {
		fmt.Println(s.Stat.Name, s.BaseStat)
	}

}
