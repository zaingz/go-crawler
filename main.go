package main

import (
	"github.com/zaingz/monzo_spyder/crawler"
	"github.com/zaingz/monzo_spyder/fetcher"
	"github.com/zaingz/monzo_spyder/linkExtractor"
	"github.com/zaingz/monzo_spyder/retry"
)

func main() {
	crawler := crawler.NewCrawler(
		"https://monzo.com",
		&fetcher.HttpFetcher{},
		&linkExtractor.HtmlLinkExtractor{},
		retry.NewSimpleDelayRetry(),
	)
	crawler.Crawl()
}
