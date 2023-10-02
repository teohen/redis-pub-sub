package main

import (
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

func main() {
	fmt.Println("vim-go")

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PORT"),
		DB:       0,
	})

}
