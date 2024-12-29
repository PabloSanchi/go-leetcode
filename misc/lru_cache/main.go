package main

import (
	"fmt"
	"lrucache/lru"
)

func main() {

	lru := lru.NewLruCache[string, int](2)

	lru.Put("a", 1)
	fmt.Println(lru)

	lru.Put("b", 2)
	fmt.Println(lru)

	lru.Put("a", 3)
	fmt.Println(lru)

	lru.Put("c", 4)
	fmt.Println(lru)

	lru.Put("a", 5)
	fmt.Println(lru)

	value, ok := lru.Get("a")
	if !ok {
		panic("Key not found")
	}

	fmt.Println(value)

}
