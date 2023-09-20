package structs

import (
	"cmp"
	"sync"
	"sync/atomic"
)

type node[K cmp.Ordered, V any] struct {
	sync.Mutex
	key    K
	item   V
	marked atomic.Bool // zero value is false
	next   atomic.Pointer[node[K, V]]
}

type List[K cmp.Ordered, V any] struct {
	head *node[K, V]
}

// type ConcurrentList[K cmp.Ordered, V any] interface {
// 	Find(K, V, bool)
// 	Insert(K, V) bool   // returns true if inserted, else false (already there)
// 	Remove(K) (V, bool) // returns val, ok (false if no node with key)
// }

func New[K cmp.Ordered, V any](minKey K, maxKey K) List[K, V] {
	tail := new(node[K, V])
	tail.key = maxKey
	head := new(node[K, V])
	head.key = minKey
	head.next.Store(tail)

	list := List[K, V]{head: head}
	return list
}

// does not remove node, else same as remove
func (l List[K, V]) Find(key K) (V, bool) {
	curr := l.head // no load??
	for curr.key < key {
		curr = curr.next.Load()
	}
	return curr.item, curr.key == key && !curr.marked.Load()
}

// returns true if inserted, else false (already there)
func (l List[K, V]) Insert(key K, val V) bool {
	// infinite loop
	for {
		// traverse without locking
		pred := l.head
		curr := pred.next.Load()

		for curr.key < key {
			pred = curr
			curr = curr.next.Load()
		}

		// lock when found
		pred.Lock()
		curr.Lock()

		// verify nodes are correct
		// if not, release locks and start over
		if !l.Validate(pred, curr) {
			curr.Unlock()
			pred.Unlock()
			continue
		}

		result := false

		if key != curr.key {
			newNode := new(node[K, V])
			newNode.key = key
			newNode.item = val
			newNode.next.Store(curr)
			pred.next.Store(newNode)
			result = true
		}

		curr.Unlock()
		pred.Unlock()
		return result
	}
}

// returns val, ok (false if no node with key)
func (l List[K, V]) Remove(key K) (V, bool) {
	for {
		pred := l.head
		curr := pred.next.Load()

		for curr.key < key {
			pred = curr
			curr = curr.next.Load()
		}

		pred.Lock()
		curr.Lock()

		if !l.Validate(pred, curr) { // failed
			curr.Unlock()
			pred.Unlock()
			continue
		}

		result := false
		if curr.key == key {
			curr.marked.Store(true)
			next := curr.next.Load()
			pred.next.Store(next)
		}

		curr.Unlock()
		pred.Unlock()

		return curr.item, result
	}
}

// pred comes before curr and curr matches
func (l List[K, V]) Validate(pred, curr *node[K, V]) bool {
	return !pred.marked.Load() && !curr.marked.Load() && pred.next.Load() == curr
}
