package list

import (
	"cmp"

	"github.com/anyaguuu/concurrent_lists/node"
)

type List[K cmp.Ordered, V any] struct {
	head *node.Node[K, V]
}
