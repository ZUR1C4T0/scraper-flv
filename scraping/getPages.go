package scraping

import (
	"strconv"

	"github.com/gocolly/colly"
)

// getPages returns the number of pages
func getPages(c *colly.Collector) int {
	var pages int

	c.OnHTML("ul.pagination li:nth-last-child(2) a", func(e *colly.HTMLElement) {
		num, _ := strconv.Atoi(e.Text)
		pages = num
	})

	c.Visit(directoryUrl)
	c.Wait()
	return pages
}
