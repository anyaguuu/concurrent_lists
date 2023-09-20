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
	var wg sync.WaitGroup
	fmt.Println("Starting")
	wg.Add(5)

	lst := structs.New[int, int](0, 10)

	go wgInsert(&wg, lst, 1, 1)
	go wgInsert(&wg, lst, 2, 2)
	go wgInsert(&wg, lst, 3, 3)
	go wgInsert(&wg, lst, 4, 4)
	go wgInsert(&wg, lst, 5, 5)

	wg.Wait()

	print(lst, 10)

}

func print[K cmp.Ordered, V any](lst structs.List[K, V], tailKey K) {
	curr := lst.Head

	for curr.Key <= tailKey {
		fmt.Println(curr.Key, ", ")
		curr = curr.Next.Load()
	}
}

func test1() {
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
