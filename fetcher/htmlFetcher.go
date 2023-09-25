package fetcher

import "net/http"

type HttpFetcher struct{}

func (f *HttpFetcher) Fetch(url string) (*http.Response, error) {
	return http.Get(url)
}
