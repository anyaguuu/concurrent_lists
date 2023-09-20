package main

import (
	"fmt"

	"github.com/anyaguuu/concurrent_lists/structs"
)

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
		fmt.Println("could not find ")
	}

}
