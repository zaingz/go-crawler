package retry

import "time"

type RetryInterface interface {
	ShouldRetry(status int) bool
	HandleRetry(url string, enqueueFunc func(string, time.Duration))
}
