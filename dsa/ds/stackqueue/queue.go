// Package queue implements a queue.
package stackqueue

import "container/list"

// Q is the internal representation of the data structure.
type Queue struct {
	l *list.List
}

// Init initializes the queue data structure.
// A queue must be initialized before it can be used.
// O(1)
func (q *Queue) Init() {
	q.l = list.New()
}

func NewQueue() *Queue {
	return &Queue{l: list.New()}
}

// Push enqueues an element to the queue.
// O(1)
func (q *Queue) Push(v interface{}) {
	q.l.PushFront(v)
}

// Pop dequeues an element from the queue.
// O(1)
func (q *Queue) Pop() interface{} {
	if q.l.Len() == 0 {
		return nil
	}

	v := q.l.Back()
	return q.l.Remove(v)
}

// Len returns the number of elements in the queue.
// O(1)
func (q *Queue) Len() int {
	return q.l.Len()
}

// IsEmpty returns true the queue has no elements.
// O(1)
func (q *Queue) IsEmpty() bool {
	return q.l.Len() == 0
}

func (q *Queue) hasKey(key string) bool {
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
