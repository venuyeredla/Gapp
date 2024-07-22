package stack_queue

import "container/list"

type Queue struct {
	l *list.List
}

func (queue *Queue) Init() {
	queue.l = list.New()
}

func NewQueue() *Queue {
	return &Queue{l: list.New()}
}

func (queue *Queue) Push(v interface{}) {
	queue.l.PushFront(v)
}

func (queue *Queue) Pop() interface{} {
	if queue.l.Len() == 0 {
		return nil
	}

	v := queue.l.Back()
	return queue.l.Remove(v)
}

func (queue *Queue) IsEmpty() bool {
	return queue.l.Len() == 0
}

func (queue *Queue) Len() int {
	return queue.l.Len()
}

func (queue *Queue) hasKey(key string) bool {
	return true
}
func (q *Queue) removeKey(key string) bool {
	return true
}

/* Queue Problems */
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
