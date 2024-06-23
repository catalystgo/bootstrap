package main

import (
	"context"
	"strconv"

	"github.com/catalystgo/bootstrap/redis"
	"github.com/containerd/containerd/pkg/randutil"
)

func main() {
	ctx := context.Background()
	addr := []string{"redis://localhost:16379", "redis://localhost:26379", "redis://localhost:36379"}

	client, err := redis.NewShardedClient(ctx, addr)
	if err != nil {
		panic(err)
	}

	defer client.Close()

	var (
		key1 = strconv.Itoa(randutil.Int())
		key2 = strconv.Itoa(randutil.Int())
		key3 = strconv.Itoa(randutil.Int())
	)

	for _, key := range []string{key1, key2, key3} {
		client, err := client.GetShard(key)
		if err != nil {
			panic(err)
		}

		val := strconv.Itoa(randutil.Int())
		if err := client.Set(ctx, key, val, 0).Err(); err != nil {
			panic(err)
		}
	}

	for _, key := range []string{key1, key2, key3} {
		client, err := client.GetShard(key)
		if err != nil {
			panic(err)
		}

		val, err := client.Get(ctx, key).Result()
		if err != nil {
			panic(err)
		}

		println(key, val, "✅")
	}
}
