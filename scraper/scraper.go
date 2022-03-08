// scraper holds all the scraping logic. All subpackage must implement ``scraper.ScrapeIterator``
package scraper

import (
	"context"
	"errors"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var Finish = errors.New("end of scrape queue")

type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

type ScrapData struct {
	Index   int
	Content *goquery.Selection
}

// ScrapeIterator is an interface to fetch data queue in the implementor.
// Implementor must return Finish error if there's no more data to fetch.
type ScrapeIterator interface {
	// Returns the next data. If there's no more data to fetch. ``scraper.Finish`` error will be returned.
	Next(ctx context.Context, url string) (data *ScrapData, err error)
}
