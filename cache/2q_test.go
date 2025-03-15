package cache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew2Q(t *testing.T) {
	q := New2Q[string, int](10)
	require.NotNil(t, q)
	require.Equal(t, 0, q.Len())
}

func TestLRUCache_Set(t *testing.T) {
	q := New2Q[string, int](10)
	q.Set("key1", 1)
	require.Equal(t, 1, q.Len())
}

func TestLRUCache_Get(t *testing.T) {
	q := New2Q[string, int](10)
	q.Set("key1", 1)
	value, found := q.Get("key1")
	require.True(t, found)
	require.Equal(t, 1, value)
}

func TestLRUCache_Del(t *testing.T) {
	q := New2Q[string, int](10)
	q.Set("key1", 1)
	q.Del("key1")
	require.Equal(t, 0, q.Len())
}

func TestLRUCache_Contains(t *testing.T) {
	q := New2Q[string, int](10)
	q.Set("key1", 1)
	require.True(t, q.Contains("key1"))
	require.False(t, q.Contains("key2"))
}

func TestLRUCache_Peek(t *testing.T) {
	q := New2Q[string, int](10)
	q.Set("key1", 1)
	value, found := q.Peek("key1")
	require.True(t, found)
	require.Equal(t, 1, value)
}

func TestLRUCache_Purge(t *testing.T) {
	q := New2Q[string, int](10)
	q.Set("key1", 1)
	q.Purge()
	require.Equal(t, 0, q.Len())
}

func TestLRUCache_Keys(t *testing.T) {
	q := New2Q[string, int](10)
	q.Set("key1", 1)
	q.Set("key2", 2)
	keys := q.Keys()
	require.Equal(t, 2, len(keys))
	require.Contains(t, keys, "key1")
	require.Contains(t, keys, "key2")
}

func TestLRUCache_Len(t *testing.T) {
	q := New2Q[string, int](10)
	q.Set("key1", 1)
	require.Equal(t, 1, q.Len())
}
