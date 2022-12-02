package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for i := 1; i <= 10; i++ {
		_, err = conn.Do("lpush", "mylist", i)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("complete to add all element")

}
