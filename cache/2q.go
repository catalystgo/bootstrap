package cache

import (
	lru "github.com/hashicorp/golang-lru"
)

type TwoQueue[K comparable, V any] struct {
	cache *lru.TwoQueueCache
}

func New2Q[K comparable, V any](capacity int) *TwoQueue[K, V] {
	cache, _ := lru.New2Q(capacity)
	return &TwoQueue[K, V]{cache: cache}
}

func (c *TwoQueue[K, V]) Set(key K, value V) {
	c.cache.Add(key, value)
}

func (c *TwoQueue[K, V]) Get(key K) (V, bool) {
	v, ok := c.cache.Get(key)
	if !ok {
		return *new(V), ok
	}
	return v.(V), ok
}

func (c *TwoQueue[K, V]) Del(key K) {
	c.cache.Remove(key)
}

func (c *TwoQueue[K, V]) Contains(key K) bool {
	return c.cache.Contains(key)
}

func (c *TwoQueue[K, V]) Peek(key K) (V, bool) {
	v, ok := c.cache.Peek(key)
	if !ok {
		return *new(V), ok
	}
	return v.(V), ok
}

func (c *TwoQueue[K, V]) Purge() {
	c.cache.Purge()
}

func (c *TwoQueue[K, V]) Keys() []K {
	keys := c.cache.Keys()
	result := make([]K, len(keys))
	for i, key := range keys {
		result[i] = key.(K)
	}
	return result
}

func (c *TwoQueue[K, V]) Len() int {
	return c.cache.Len()
}
