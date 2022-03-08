// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	scraper "github.com/tigorlazuardi/light-novel-translation-scraper/scraper"
)

// ScrapeIterator is an autogenerated mock type for the ScrapeIterator type
type ScrapeIterator struct {
	mock.Mock
}

// Next provides a mock function with given fields: ctx, url
func (_m *ScrapeIterator) Next(ctx context.Context, url string) (*scraper.ScrapData, error) {
	ret := _m.Called(ctx, url)

	var r0 *scraper.ScrapData
	if rf, ok := ret.Get(0).(func(context.Context, string) *scraper.ScrapData); ok {
		r0 = rf(ctx, url)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scraper.ScrapData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
