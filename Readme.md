
# Web Crawler
This is a robust and modular simple web crawler designed to efficiently traverse and fetch web pages. Built with a focus on extensibility, configurability, and best coding practices.

## Features

- **Modular Design**: Clear separation of concerns with distinct modules for fetching, link extraction, retry logic, and queue management.
- **Concurrency Control**: Efficiently manages multiple concurrent requests without overwhelming target servers.
- **URL Deduplication**: Ensures each URL is processed only once for optimal performance.
- **Retry Mechanism**: Handles transient failures by re-enqueuing URLs with a delay.
- **Configurable**: Easily swap out core components to suit different needs, thanks to the use of interfaces and the Strategy Pattern.
- **Robustness**: Designed to handle edge cases and the unpredictable nature of the web.

## Design Patterns

- **Single Responsibility Principle**: Each module has a distinct responsibility, making the code maintainable and testable.
- **Strategy Pattern**: Use of interfaces allows for runtime selection of algorithms, making the crawler highly configurable.
- **Observer Pattern**: The crawler observes all the URLs fetches and delegates retry logic, ensuring resilience.

## Modules

1. **Fetcher**: Handles HTTP requests.
2. **Link Extractor**: Parses fetched HTML content to extract links.
3. **Retry Manager**: Implements retry logic for failed requests.
4. **Queue**: Manages the list of URLs to be processed, ensuring concurrency control and deduplication.

## Usage

```go
crawler := NewCrawler(
    startURL,
    fetcherInstance,
    linkExtractorInstance,
    retryManagerInstance,
)
crawler.Crawl()
```

Replace `startURL`, `fetcherInstance`, `linkExtractorInstance`, and `retryManagerInstance` with appropriate values.

## Extending the Crawler

Because of its modular design, extending the crawler is straightforward:

- **New Fetch Mechanisms**: Implement the `FetcherInterface` to add new fetch mechanisms.
- **Link Extraction Techniques**: Implement the `LinkExtractorInterface` for different link extraction strategies.
- **Custom Retry Logic**: Implement the `RetryInterface` to define custom retry behaviors.

## Conclusion

This web crawler is simple but yet versatile tool designed with a balance of efficiency, extensibility, and robustness. 