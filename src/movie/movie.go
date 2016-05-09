package main

import (
	"fmt"
	"encoding/json"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "AAA", Year: 1999, Color: false,
		Actors: []string{"HM", "IB"}},
	{Title: "BBB", Year: 1111, Color: true,
		Actors: []string{"PN"}},
}

func main() {
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marsha1ling failed: %s", err)
	}

	fmt.Printf("%s", data)
}