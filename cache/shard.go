package cache

import (
	"strconv"
	"sync"

	"github.com/serialx/hashring"
)

type Cache[K comparable, V any] interface {
	Set(key K, value V)
	Get(key K) (V, bool)
	Del(key K)
	Contains(key K) bool
	Peek(key K) (V, bool)
	Purge()
	Keys() []K
	Len() int
}

type ShardedCache[K string, V any] struct {
	ring   *hashring.HashRing
	shards []Cache[K, V]
	mu     sync.RWMutex
}

func NewShardedCache[K string, V any](
	numShards int,
	sizePerShard int,
	cacheFactory func(size int) Cache[K, V],
) *ShardedCache[K, V] {
	shards := make([]Cache[K, V], numShards)
	for i := 0; i < numShards; i++ {
		shards[i] = cacheFactory(sizePerShard)
	}

	nodes := make([]string, numShards)
	for i := 0; i < numShards; i++ {
		nodes[i] = string(rune(i))
	}

	return &ShardedCache[K, V]{
		ring:   hashring.New(nodes),
		shards: shards,
	}
}

func (sc *ShardedCache[K, V]) getShard(key K) Cache[K, V] {
	sc.mu.RLock()
	defer sc.mu.RUnlock()
	node, _ := sc.ring.GetNode(string(key))
	index, _ := strconv.Atoi(node)
	return sc.shards[index]
}

func (sc *ShardedCache[K, V]) Set(key K, value V) {
	shard := sc.getShard(key)
	shard.Set(key, value)
}

func (sc *ShardedCache[K, V]) Get(key K) (V, bool) {
	shard := sc.getShard(key)
	return shard.Get(key)
}

func (sc *ShardedCache[K, V]) Del(key K) {
	shard := sc.getShard(key)
	shard.Del(key)
}

func (sc *ShardedCache[K, V]) Len() int {
	sc.mu.RLock()
	defer sc.mu.RUnlock()
	total := 0
	for _, shard := range sc.shards {
		total += shard.Len()
	}
	return total
}
