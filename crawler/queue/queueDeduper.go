package crawler

import "sync"

type urlQueueDeduper struct {
	seen  map[string]bool
	mutex sync.Mutex
}

func newUrlQueueDeduper() *urlQueueDeduper {
	return &urlQueueDeduper{seen: make(map[string]bool)}
}

func (d *urlQueueDeduper) isSeen(url string) bool {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	_, exists := d.seen[url]
	return exists
}

func (d *urlQueueDeduper) markAsSeen(url string) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.seen[url] = true
}

func (d *urlQueueDeduper) countSeen() int {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	return len(d.seen)
}
