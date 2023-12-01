package scraping

import (
	"github.com/ZUR1C4T0/scraper-flv/models/anime"
	"github.com/gocolly/colly"
)

// ScrapeAnimes scrapes the animes from the website
func ScrapeAnimes() []anime.Anime {
	c := colly.NewCollector(colly.Async(true))

	print("Pages...\r")
	pages := getPages(c)
	println("Pages: ", pages)

	print("Links...\r")
	links := getAnimeLinks(c, pages)
	println("Links: ", len(links))

	print("Animes...\r")
	animes := getAnimes(c, links)
	println("Animes: ", len(animes))

	return animes
}
