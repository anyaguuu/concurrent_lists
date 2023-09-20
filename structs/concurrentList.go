package structs

import (
	"cmp"
)

// type ConcurrentList[K cmp.Ordered, V any] interface {
// 	Find(K, V, bool)    // does not remove node, else same as remove
// 	Insert(K, V) bool   // returns true if inserted, else false (already there)
// 	Remove(K) (V, bool) // returns val, ok (false if no node with key)
// }

func New[K cmp.Ordered, V any](minKey K, maxKey K) List[K, V] {

}
