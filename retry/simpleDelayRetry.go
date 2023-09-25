package retry

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

const (
	MaxDelay   = 1 * time.Second
	MaxRetries = 2
)

var retryStatuses = []int{
	http.StatusTooManyRequests,
	http.StatusServiceUnavailable,
	http.StatusInternalServerError,
	http.StatusGatewayTimeout,
	http.StatusNotFound,
}

type SimpleDelayRetry struct {
	retryCounts map[string]int
	mutex       sync.Mutex
}

func NewSimpleDelayRetry() *SimpleDelayRetry {
	return &SimpleDelayRetry{
		retryCounts: make(map[string]int),
	}
}

func (r *SimpleDelayRetry) ShouldRetry(status int) bool {
	for _, s := range retryStatuses {
		if s == status {
			return true
		}
	}
	return false
}

func (r *SimpleDelayRetry) incrementAndCheckRetryCount(url string) bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.retryCounts[url]++
	if r.retryCounts[url] > MaxRetries {
		fmt.Printf("Max retries reached for URL: %s\n", url)
		return false
	}
	return true
}

func (r *SimpleDelayRetry) getRetryCount(url string) int {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.retryCounts[url]
}

func (r *SimpleDelayRetry) HandleRetry(url string, enqueueFunc func(string, time.Duration)) {

	if !r.incrementAndCheckRetryCount(url) {
		return
	}

	retryCount := r.getRetryCount(url)
	delay := time.Duration(retryCount)*time.Second + time.Duration(rand.Int63n(int64(MaxDelay)))
	fmt.Printf("Retrying in (%v) for URL (Attempt %d): %s\n", delay, retryCount, url)
	enqueueFunc(url, delay)
}
