package main

import (
	"fmt"

	"github.com/catalystgo/bootstrap/cache"
)

func main() {
	// Create a sharded cache with 4 shards, each with a capacity of 100
	shardedCache := cache.NewShardedCache[string, string](4, 100, func(size int) cache.Cache[string, string] {
		return cache.New2Q[string, string](size)
	})

	// Set some values
	shardedCache.Set("key1", "value1")
	shardedCache.Set("key2", "value2")
	shardedCache.Set("key3", "value3")

	// Get and print values
	if value, ok := shardedCache.Get("key1"); ok {
		fmt.Println("key1:", value)
	}
	if value, ok := shardedCache.Get("key2"); ok {
		fmt.Println("key2:", value)
	}
	if value, ok := shardedCache.Get("key3"); ok {
		fmt.Println("key3:", value)
	}

	// Print the total number of items in the cache
	fmt.Println("Total items in cache:", shardedCache.Len())
}
