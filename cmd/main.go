package main

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

func main() {
	fmt.Println("vim-go")

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PORT"),
		DB:       0,
	})

	fmt.Println(rdb)

	//scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Say something")

	/*for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			break
		}

		commands := strings.Split(text, " ")

		if commands[0] == "sub" {

		}

		fmt.Println("You entered: ", text)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error", err)
	}
	*/
	pubsub := rdb.Subscribe("jose")

	ch := pubsub.Channel()

	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}

}
