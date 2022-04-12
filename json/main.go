package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Movie struct {
	Title   string
	Year    string
	Rated   string
	Resumo  string `json:"Plot"`
	Ratings []Rating
}

type Rating struct {
	Source string
	Value  string
}

func main() {
	words := []string{"The", "Batman"}
	movie, err := Search(words)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	fmt.Printf("Movie: %+v", *movie)

}

func Search(words []string) (*Movie, error) {
	q := url.QueryEscape(strings.Join(words, " "))
	resp, err := http.Get("http://www.omdbapi.com/?apikey=8b9a8b71&plot=full&t=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Search query failed: %s", resp.Status)
	}

	var movie Movie
	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &movie, nil
}
