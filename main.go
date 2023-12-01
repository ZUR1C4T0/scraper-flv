package main

import (
	"fmt"

	"github.com/ZUR1C4T0/scraper-flv/scraping"
	"github.com/ZUR1C4T0/scraper-flv/storage"
)

func main() {
	animeList := scraping.ScrapeAnimes()

	error := storage.JsonStorage("./animes.json", animeList)
	if error != nil {
		fmt.Println(error)
	}
}
