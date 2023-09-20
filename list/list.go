package list

import "cmp"

type List[K cmp.Ordered, V any] struct {
	head *node.node[K, V]
}
