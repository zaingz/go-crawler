package crawler

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	crawler "github.com/zaingz/monzo_spyder/crawler/queue"
	"github.com/zaingz/monzo_spyder/fetcher"
	"github.com/zaingz/monzo_spyder/linkExtractor"
	"github.com/zaingz/monzo_spyder/retry"
)

const MaxConcurrentRequests = 100

type Crawler struct {
	startURL      string
	fetcher       fetcher.FetcherInterface
	linkExtractor linkExtractor.LinkExtractorInterface
	startTime     time.Time
	queue         *crawler.UrlQueue
	retryManager  retry.RetryInterface
}

func NewCrawler(startURL string, fetcher fetcher.FetcherInterface, linkExtractor linkExtractor.LinkExtractorInterface, rm retry.RetryInterface) *Crawler {
	return &Crawler{
		startURL:      startURL,
		fetcher:       fetcher,
		linkExtractor: linkExtractor,
		queue:         crawler.NewUrlQueue(),
		retryManager:  rm,
	}
}

func (c *Crawler) isSameDomain(link string) bool {
	baseURL, err := url.Parse(c.startURL)
	if err != nil {
		return false
	}

	linkURL, err := url.Parse(link)
	if err != nil {
		return false
	}

	return baseURL.Hostname() == linkURL.Hostname()
}

func (c *Crawler) Crawl() {
	c.startTime = time.Now()
	semaphore := make(chan struct{}, MaxConcurrentRequests)

	c.queue.Enqueue(c.startURL, false)

	go func() {
		for url := range c.queue.Urls {
			semaphore <- struct{}{}
			go func(u string) {
				c.processURL(u)
				<-semaphore
			}(url)
		}
	}()

	c.queue.WaitAndClose()
	printSummary(c.queue.UniqueUrls(), c.startTime)
}

func (c *Crawler) processURL(url string) {
	defer c.queue.Done()

	resp, err := c.fetcher.Fetch(url)
	if err != nil {
		fmt.Printf("Error fetching URL %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	if c.retryManager.ShouldRetry(resp.StatusCode) {
		c.retryManager.HandleRetry(url, c.queue.RetryHandler.EnqueueAfterDelay)
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error HTTP status: %s URL: %s\n", resp.Status, url)
		return
	}

	fmt.Println("Visited: ", url)

	links, err := c.linkExtractor.ExtractLinks(resp)
	if err != nil {
		fmt.Printf("Error extracting links from %s: %v\n", url, err)
		return
	}

	for _, link := range links {
		if c.isSameDomain(link) {
			c.queue.Enqueue(link, false)
		}
	}
}
