package main

import (
	"context"

	"github.com/catalystgo/bootstrap/redis"
)

func main() {
	ctx := context.Background()

	client, err := redis.NewClient(ctx, "redis://localhost:6379")
	if err != nil {
		panic(err)
	}

	defer client.Close()

	if _, err := client.Ping(ctx).Result(); err != nil {
		panic(err)
	}

	if err := client.Set(ctx, "mykey", "Hello, World!", 0).Err(); err != nil {
		panic(err)
	}

	val, err := client.Get(ctx, "mykey").Result()
	if err != nil {
		panic(err)
	}

	println(val, "✅")
}
