package concurrentlist

import "cmp"

type ConcurrentList[K cmp.Ordered, V any] interface {
	Find(K, V, bool)
	Insert(K, V) bool
	Remove(K) (V, bool)
}
