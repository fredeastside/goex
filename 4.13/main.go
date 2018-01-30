package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

type Movie struct {
	Title string
	Year  string
	Error string
}

func main() {
	var title string
	flag.StringVar(&title, "t", "", "Title of movie for search")
	flag.Parse()
	response, err := http.Get("http://www.omdbapi.com/?apikey=73f5a642&t=" + title)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	defer response.Body.Close()

	var movie Movie
	if err := json.NewDecoder(response.Body).Decode(&movie); err != nil {
		fmt.Printf("Error: %v", err)
	}
	if movie.Error != "" {
		fmt.Printf("Error: %s", movie.Error)
		return
	}
	fmt.Printf("Movie Title: %s\n", movie.Title)
	fmt.Printf("Movie Year: %s\n", movie.Year)
}
