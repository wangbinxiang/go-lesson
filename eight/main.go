package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var client redis.UniversalClient

func main() {
	client = redis.NewClient(&redis.Options{
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
	client.FlushAll(ctx)
	res, err := client.Info(ctx, "memory").Result()
	if err != nil {
		fmt.Println(err)
	} else {
		memLen := strings.Split(res, "\r\n")[2]
		fmt.Printf("base %s\n", memLen)
	}

	fmt.Println("reis 连接成功：", pong)
	writeData(1000, 10, "len10_10k")
	writeData(5000, 10, "len10_50k")
	writeData(50000, 10, "len10_500k")

	writeData(1000, 1000, "len1000_10k")
	writeData(50000, 1000, "len1000_50k")
	writeData(500000, 1000, "len1000_500k")

	writeData(10000, 5000, "len5000_10k")
	writeData(50000, 5000, "len5000_50k")
	writeData(500000, 5000, "len5000_500k")
}

func writeData(num, size int, key string) {
	client.FlushAll(ctx)
	arr := make([]byte, size)
	for i := 0; i < size; i++ {
		arr[i] = 'x'
	}
	data := string(arr)

	for i := 0; i < num; i++ {
		k := fmt.Sprintf("%s:%d", key, i)
		cmd := client.Set(ctx, k, data, -1)
		err := cmd.Err()
		if err != nil {
			fmt.Println(cmd.String())
		}
	}

	res, err := client.Info(ctx, "memory").Result()
	if err != nil {
		fmt.Println(err)
	} else {
		memLen := strings.Split(res, "\r\n")[2]
		fmt.Printf("%s %s\n", key, memLen)
	}
}
