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
// 	Find(K, V, bool)    // does not remove node, else same as remove
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
