package main

import (
	"cmp"
	"fmt"
	"sync"

	"github.com/anyaguuu/concurrent_lists/structs"
)

func wgFind[K cmp.Ordered, V any](wg *sync.WaitGroup, lst structs.List[K, V], key K) (V, bool) {
	defer wg.Done()
	return lst.Find(key)
}
func wgInsert[K cmp.Ordered, V any](wg *sync.WaitGroup, lst structs.List[K, V], key K, val V) bool {
	defer wg.Done()
	return lst.Insert(key, val)
}
func wgRemove[K cmp.Ordered, V any](wg *sync.WaitGroup, lst structs.List[K, V], key K) (V, bool) {
	defer wg.Done()
	return lst.Remove(key)
}

func main() {
	lst := structs.New[int, int](0, 10)
	lst.Insert(1, 1)
	lst.Insert(2, 2)
	lst.Insert(3, 3)
	lst.Insert(10, 10)
	lst.Insert(4, 4)

	find := 3
	_, ok := lst.Find(find)
	if ok {
		fmt.Println("found", find)
	} else {
		fmt.Println("could not find", find)
	}

	remove := 4
	val, ok := lst.Remove(remove)
	if ok {
		fmt.Println("removed", val)
	} else {
		fmt.Println("could not remove", remove)
	}

}
