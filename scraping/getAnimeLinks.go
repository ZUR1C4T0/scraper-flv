package scraping

import (
	"fmt"

	"github.com/gocolly/colly"
)

func getAnimeLinks(col *colly.Collector, pages int) []string {
	c := col.Clone()
	arraySize := 24 * pages
	links := make([]string, 0, arraySize)

	c.OnHTML(".ListAnimes", func(e *colly.HTMLElement) {
		e.ForEach(".Anime > a", func(_ int, h *colly.HTMLElement) {
			link := baseUrl + h.Attr("href")
			links = append(links, link)
		})
	})

	for i := 1; i <= pages; i++ {
		go c.Visit(fmt.Sprintf("%s?page=%d", directoryUrl, i))
	}

	c.Wait()
	return links
}
