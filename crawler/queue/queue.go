package crawler

import (
	"sync"
)

type UrlQueue struct {
	Urls         chan string
	wg           sync.WaitGroup
	deduper      *urlQueueDeduper
	RetryHandler *urlQueueRetryHandler
}

func NewUrlQueue() *UrlQueue {
	q := &UrlQueue{
		Urls:    make(chan string, 1000),
		deduper: newUrlQueueDeduper(),
	}
	q.RetryHandler = q.newUrlQueueRetryHandler()
	return q
}

func (q *UrlQueue) Enqueue(url string, bypassDedupe bool) {
	if !bypassDedupe && q.deduper.isSeen(url) {
		return
	}
	q.deduper.markAsSeen(url)

	q.wg.Add(1)
	go func() {
		q.Urls <- url
	}()
}

func (q *UrlQueue) Done() {
	q.wg.Done()
}

func (q *UrlQueue) UniqueUrls() int {
	return q.deduper.countSeen()
}

func (q *UrlQueue) WaitAndClose() {
	q.wg.Wait()
	close(q.Urls)
}
