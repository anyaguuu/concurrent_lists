package structs

import (
	"cmp"
	"sync"
	"sync/atomic"
)

type Node[K cmp.Ordered, V any] struct {
	sync.Mutex
	key    K
	item   V
	marked atomic.Bool
	next   atomic.Pointer[Node[K, V]]
}
