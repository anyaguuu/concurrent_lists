package structs

import (
	"cmp"
	"sync"
	"sync/atomic"
)

type node[K cmp.Ordered, V any] struct { // comment
	sync.Mutex
	Key    K
	item   V
	marked atomic.Bool // zero value is false
	Next   atomic.Pointer[node[K, V]]
}

type List[K cmp.Ordered, V any] struct {
	Head *node[K, V]
}

// type ConcurrentList[K cmp.Ordered, V any] interface {
// 	Find(K, V, bool)
// 	Insert(K, V) bool   // returns true if inserted, else false (already there)
// 	Remove(K) (V, bool) // returns val, ok (false if no node with key)
// }

func New[K cmp.Ordered, V any](minKey K, maxKey K) List[K, V] {
	tail := new(node[K, V])
	tail.Key = maxKey
	head := new(node[K, V])
	head.Key = minKey
	head.Next.Store(tail)

	list := List[K, V]{Head: head}
	return list
}

// does not remove node, else same as remove
func (l List[K, V]) Find(key K) (V, bool) {
	curr := l.Head // no load??
	for curr.Key < key {
		curr = curr.Next.Load()
	}
	return curr.item, curr.Key == key && !curr.marked.Load()
}

// returns true if inserted, else false (already there)
func (l List[K, V]) Insert(key K, val V) bool {
	// infinite loop
	for {
		// traverse without locking
		pred := l.Head
		curr := pred.Next.Load()

		for curr.Key < key {
			pred = curr
			curr = curr.Next.Load()
		}

		// lock when found
		pred.Lock()
		curr.Lock()

		// verify nodes are correct
		// if not, release locks and start over
		if !l.Validate(pred, curr) {
			curr.Unlock()
			pred.Unlock()
			return false
		} else {
			newNode := new(node[K, V])
			newNode.Key = key
			newNode.item = val
			newNode.Next.Store(curr)
			pred.Next.Store(newNode)
			curr.Unlock()
			pred.Unlock()
			return true
		}
	}
}

// returns val, ok (false if no node with key)
func (l List[K, V]) Remove(key K) (V, bool) {
	for {
		pred := l.Head
		curr := pred.Next.Load()

		for curr.Key < key {
			pred = curr
			curr = curr.Next.Load()
		}

		pred.Lock()
		curr.Lock()

		if l.Validate(pred, curr) { // failed
			if curr.Key == key {
				curr.marked.Store(true)
				Next := curr.Next.Load()
				pred.Next.Store(Next)
				return curr.item, true
			}

			curr.Unlock()
			pred.Unlock()
		} else {
			curr.Unlock()
			pred.Unlock()
			return l.Head.item, false
		}
	}
}

// pred comes before curr and curr matches
func (l List[K, V]) Validate(pred, curr *node[K, V]) bool {
	return !pred.marked.Load() && !curr.marked.Load() && pred.Next.Load() == curr
}
