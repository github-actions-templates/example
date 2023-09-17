package main

import (
	"context"
	"time"

	redis "github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	if err := rdb.Set(ctx, "test", time.Now().String(), time.Second*10).Err(); err != nil {
		panic(err)
	}

	if result := rdb.Get(ctx, "test"); result.Err() != nil {
		panic(result.Err())
	} else {
		println(result.Val())
	}
}
