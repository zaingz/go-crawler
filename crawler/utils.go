package crawler

import (
	"fmt"
	"strings"
	"time"
)

func printSummary(urlsVisited int, startTime time.Time) {
	duration := time.Since(startTime)
	urlsPerSecond := float64(urlsVisited) / duration.Seconds()
	avgTimePerURL := duration / time.Duration(urlsVisited)
	separator := strings.Repeat("-", 40)

	fmt.Println()
	fmt.Println(separator)
	fmt.Println("  Crawling Summary")
	fmt.Println(separator)
	fmt.Printf("  Total URLs visited:    %d\n", urlsVisited)
	fmt.Printf("  Time taken:            %v\n", duration)
	fmt.Printf("  Average time per URL:  %v\n", avgTimePerURL)
	fmt.Printf("  URLs visited/second:   %.2f\n", urlsPerSecond)
	fmt.Println(separator)
}
