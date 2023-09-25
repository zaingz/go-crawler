package fetcher

import "net/http"

type FetcherInterface interface {
	Fetch(url string) (*http.Response, error)
}
