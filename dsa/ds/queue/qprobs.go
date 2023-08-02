package queue

type LRUCache struct {
	queue *Queue
	Len   uint8
}

func NewLRUCache(len uint8) *LRUCache {
	q := new(Queue)
	return &LRUCache{queue: q, Len: len}
}

func (lru *LRUCache) getPage(key string) {
	if lru.queue.hasKey(key) {
		lru.queue.removeKey(key)
	} else {
		if lru.queue.Len() == lru.queue.Len() {
			lru.queue.Pop()
		}
	}
	lru.queue.Push(key)

}
