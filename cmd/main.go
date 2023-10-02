package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/go-redis/redis"
)

func subs(client redis.Client, channelName string) {
	fmt.Println("subing...")
	pubsub := client.Subscribe(channelName)

	ch := pubsub.Channel()

	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}

}

func pub(client redis.Client, channelName string, msg string) {
	fmt.Println("publishing...")
	err := client.Publish(channelName, msg).Err()

	if err != nil {
		fmt.Println("ERROR")
		panic(err)
	}

}

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PORT"),
		DB:       0,
	})

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Say something")

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			break
		}

		commands := strings.Split(text, " ")

		if commands[0] == "sub" {
			subs(*rdb, commands[1])
		}

		if commands[0] == "pub" {
			pub(*rdb, commands[1], commands[2])
		}

		fmt.Println("You entered: ", text)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error", err)
	}

}
