package lightnoveltranslation

import (
	"context"
	"time"

	"github.com/tigorlazuardi/light-novel-translation-scraper/scraper"
)

type dataTuple struct {
	data *scraper.ScrapData
	err  error
}

type Scraper struct {
	client scraper.Doer
	queue  chan dataTuple
	url    string
}

func NewScraper(client scraper.Doer, url string, queueSize int) *Scraper {
	return &Scraper{
		client: client,
		queue:  make(chan dataTuple, queueSize),
		url:    url,
	}
}

// Returns the next data. If there's no more data to fetch. ``scraper.Finish`` error will be returned.
func (s *Scraper) Next(ctx context.Context) (data *scraper.ScrapData, err error) {
	select {
	case <-ctx.Done():
		if t, ok := ctx.Deadline(); ok && time.Now().After(t) {
			return nil, context.DeadlineExceeded
		}
		return nil, context.Canceled
	case tuple, ok := <-s.queue:
		if !ok {
			return nil, scraper.Finish
		}
		return tuple.data, tuple.err
	}
}
