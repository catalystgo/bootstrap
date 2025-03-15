package cache_test

import (
	"testing"

	"github.com/catalystgo/bootstrap/cache"
	"github.com/stretchr/testify/require"
)

func cacheFactory(size int) cache.Cache[string, int] {
	return cache.New2Q[string, int](size)
}

func TestShardedCache_Set(t *testing.T) {
	shardedCache := cache.NewShardedCache[string, int](1, 100, cacheFactory)

	shardedCache.Set("key1", 1)
}

func TestShardedCache_Get(t *testing.T) {
	shardedCache := cache.NewShardedCache[string, int](1, 100, cacheFactory)

	shardedCache.Set("key1", 1)
	value, found := shardedCache.Get("key1")
	require.True(t, found)
	require.Equal(t, 1, value)
}

func TestShardedCache_Del(t *testing.T) {
	shardedCache := cache.NewShardedCache[string, int](1, 100, cacheFactory)

	shardedCache.Set("key1", 1)
	shardedCache.Del("key1")
	_, found := shardedCache.Get("key1")
	require.False(t, found)
}

func TestShardedCache_Len(t *testing.T) {
	shardedCache := cache.NewShardedCache[string, int](3, 100, cacheFactory)

	shardedCache.Set("key1", 1)
	shardedCache.Set("key2", 2)
	shardedCache.Set("key3", 3)
	require.Equal(t, 3, shardedCache.Len())
}
