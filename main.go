package main

import (
	"fmt"
	"pintu/lfu"
	"pintu/lru"
	"pintu/none"
)


func problem1() {
	limit := 3
	evictionManager := none.EvictionManager{}
	cache := none.NewInMemoryCache(limit, evictionManager)

	fmt.Println(cache.Add("key1", "value1"))   // return 0
	fmt.Println(cache.Add("key2", "value2"))   // return 0
	fmt.Println(cache.Add("key3", "value3"))   // return 0
	fmt.Println(cache.Add("key2", "value2.1")) // return 1
	fmt.Println(cache.Get("key3"))             // return value3
	fmt.Println(cache.Get("key1"))             // return value1
	fmt.Println(cache.Get("key3"))             // return value3
	fmt.Println(cache.Keys())                  // return ["key1", "key2", "key3"]
	fmt.Println(cache.Add("key4", ""))         // return / throw Error("key_limit_exceeded")
	fmt.Println(cache.Keys())                  // return ["key1", "key2", "key3"]
	fmt.Println(cache.Clear())                 // return 3
	fmt.Println(cache.Keys())                  // return []
}

func problem2() {
	limit := 3
	em := lru.EvictionManager{}

	cache := lru.NewInMemoryCache(limit, em)
	
	fmt.Println(cache.Add("key1", "value1")) // return 0
	fmt.Println(cache.Add("key2", "value2")) // return 0
	fmt.Println(cache.Add("key3", "value3")) // return 0
	fmt.Println(cache.Add("key2", "value2.1")) // return 1
	fmt.Println(cache.Get("key3")) // return value3
	fmt.Println(cache.Get("key1")) // return value1
	fmt.Println(cache.Get("key2")) // return value2.1
	fmt.Println(cache.Keys()) // return ["key1", "key2", "key3"]
	fmt.Println(cache.Add("key4", "")) // return 0
	fmt.Println(cache.Keys()) // return ["key1", "key2", "key4"]
	//(key 3 is the least recently used key, so when key4 added, we will remove key3 from cache)
	fmt.Println(cache.Clear()) // return 3
	fmt.Println(cache.Keys()) // return []
}

func problem3() {
	limit := 3
	em := lfu.EvictionManager{}

	cache := lfu.NewInMemoryCache(limit, em)

	fmt.Println(cache.Add("key1", "value1")) // return 0
	fmt.Println(cache.Add("key2", "value2")) // return 0
	fmt.Println(cache.Add("key3", "value3")) // return 0
	fmt.Println(cache.Add("key2", "value2.1")) // return 1
	fmt.Println(cache.Get("key3")) // return value3
	fmt.Println(cache.Get("key1")) // return value1
	fmt.Println(cache.Get("key2")) // return value2.1
	fmt.Println(cache.Get("key3")) // return value3
	fmt.Println(cache.Get("key1")) // return value1
	fmt.Println(cache.Keys()) // return ["key1", "key2", "key3"]
	fmt.Println(cache.Add("key4", "")) // return 0
	fmt.Println(cache.Keys()) // return ["key1", "key3", "key4"]
	// (key1 has 2 freq, key 2 has 1 freq, and key 3 has 2 freq, so when key4 added, we will remove key 2 from cache)
	fmt.Println(cache.Clear()) // return 3
	fmt.Println(cache.Keys()) // return []
}

func main() {
	/**
	Assumption

	Since I made it in Golang, I assume when add "key4" it has empty value. So, I modified
	the case to give "" (empty string) as value.

	I use slice to store data, since it concern about the order of the key,
	and I assume it won't used to store many data.

	I add private function to help me read the code easier.
	*/

	fmt.Println("Problem 1")
	problem1()

	fmt.Println("\nProblem 2")
	problem2()

	fmt.Println("\nProblem 3")
	problem3()
}

