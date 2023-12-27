package scraping

import (
	"fmt"

	"github.com/ZUR1C4T0/scraper-flv/logging"
	"github.com/ZUR1C4T0/scraper-flv/models/anime"
	"github.com/gocolly/colly"
)

func getAnimes(c *colly.Collector, links []string) []anime.Anime {
	animes := make([]anime.Anime, 0, len(links))

	c.OnError(func(r *colly.Response, err error) {
		errMsg := fmt.Sprintf("Error %d %s: %s", r.StatusCode, err.Error(), r.Request.URL.String())
		logging.LogError("anime-errors.log", errMsg)
	})

	c.OnHTML(".Body", func(e *colly.HTMLElement) {
		a := anime.Anime{}

		a.Title = e.ChildText("h1.Title")
		a.Type = anime.Type(e.ChildText(".Type"))
		e.ForEach(".TxtAlt", func(_ int, h *colly.HTMLElement) {
			a.TxtAlt = append(a.TxtAlt, h.Text)
		})
		e.ForEach("nav.Nvgnrs a", func(_ int, h *colly.HTMLElement) {
			a.Genres = append(a.Genres, h.Text)
		})
		a.Synopsis = e.ChildText(".Description p")
		a.Cover = baseUrl + e.ChildAttr(".Image img", "src")
		a.Status = anime.Status(e.ChildText(".AnmStts"))
		e.ForEach(".ListAnmRel li a", func(_ int, h *colly.HTMLElement) {
			a.RelatedAnimes = append(a.RelatedAnimes, anime.RelatedAnime{
				Title:    h.Text,
				Relation: anime.Relation(e.ChildText(".ListAnmRel li")),
			})
		})

		animes = append(animes, a)
	})

	// divide links in chunks to avoid errors of too many requests
	chunks := chunkLinks(links, 200)
	for _, chunk := range chunks {

		for _, link := range chunk {
			go c.Visit(link)
		}

		c.Wait()
	}

	return animes
}
