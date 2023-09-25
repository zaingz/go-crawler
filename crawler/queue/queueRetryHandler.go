package crawler

import (
	"sync"
	"time"
)

type urlQueueRetryHandler struct {
	enqueueFunc func(string, bool)
	wg          *sync.WaitGroup
}

func (q *UrlQueue) newUrlQueueRetryHandler() *urlQueueRetryHandler {
	return &urlQueueRetryHandler{
		enqueueFunc: q.Enqueue,
		wg:          &q.wg,
	}
}

func (r *urlQueueRetryHandler) EnqueueAfterDelay(url string, delay time.Duration) {
	r.wg.Add(1)
	go func() {
		defer r.wg.Done()
		time.Sleep(delay)
		r.enqueueFunc(url, true)
	}()

}
