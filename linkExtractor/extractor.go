package linkExtractor

import "net/http"

type LinkExtractorInterface interface {
	ExtractLinks(resp *http.Response) ([]string, error)
}
