package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		Password:     "", // no password set
		DB:           0,  // use default DB
		PoolSize:     128,
		MinIdleConns: 100,
		MaxRetries:   5,
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println("reis 连接失败：", pong, err)
		return
	}
	// client.FlushAll(ctx)
	fmt.Println("reis 连接成功：", pong)
	err = client.Set(ctx, "used_memory_human", "used_memory", 0).Err()
	fmt.Println(err)
	val, _ := client.Get(ctx, "used_memory_human").Result()
	fmt.Println(val)
	c := client.Info(ctx, "memory")
	r, _ := c.Result()
	fmt.Println(r)
	rl := strings.Split(r, "\r\n")
	fmt.Println(rl[2])
	// fmt.Println(c.Result())
	// client.FlushAll(ctx)
	// c = client.Info(ctx, "memory")
	// fmt.Println(c.Result())
}
