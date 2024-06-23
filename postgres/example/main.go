package main

import (
	"context"

	"github.com/catalystgo/bootstrap/postgres"
)

func main() {
	ctx := context.Background()

	db, err := postgres.New(ctx, `postgres://postgres:postgres@localhost:5432/bootstrap`,
		// example of setting options
		postgres.WithMaxConns(10),
		postgres.WithMinConns(5),
		postgres.WithMaxConnLifetime(5),
		postgres.WithMaxConnIdleTime(5),
	)

	if err != nil {
		panic(err)
	}

	err = db.Ping(ctx)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	println("Connected to Postgres ✅")
}
