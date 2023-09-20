package node

import (
	"cmp"
	"sync"
	"sync/atomic"
)

type node[K cmp.Ordered, V any] struct {
	sync.Mutex
	key    K
	item   V
	marked atomic.Bool
	next   atomic.Pointer[node[K, V]]
}
