package main

import (
	"fmt"
	"log"

	// Import the redigo/redis package.
	"github.com/gomodule/redigo/redis"
)

func main() {
	// Establish a connection to the Redis server listening on port
	// 6379 of the local machine. 6379 is the default port, so unless
	// you've already changed the Redis configuration file this should
	// work.
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	// Importantly, use defer to ensure the connection is always
	// properly closed before exiting the main() function.
	defer conn.Close()

	// Send our command across the connection. The first parameter to
	// Do() is always the name of the Redis command (in this example
	// HMSET), optionally followed by any necessary arguments (in this
	// example the key, followed by the various hash fields and values).
	_, err = conn.Do("HMSET", "album:1", "title", "Electric Ladyland", "artist", "Jimi Hendrix", "price", 4.95, "likes", 8)
	if err != nil {
		log.Fatal(err)
	}

	// Success to store data
	fmt.Println("Electric Ladyland added!")

	// Issue a HGET command to retrieve the title for a specific album,
	// and use the Str() helper method to convert the reply to a string.
	title, err := redis.String(conn.Do("HGET", "album:1", "title"))
	if err != nil {
		log.Fatal(err)
	}

	// Similarly, get the artist and convert it to a string.
	artist, err := redis.String(conn.Do("HGET", "album:1", "artist"))
	if err != nil {
		log.Fatal(err)
	}

	// And the price as a float64...
	price, err := redis.Float64(conn.Do("HGET", "album:1", "price"))
	if err != nil {
		log.Fatal(err)
	}

	// And the number of likes as an integer.
	likes, err := redis.Int(conn.Do("HGET", "album:1", "likes"))
	if err != nil {
		log.Fatal(err)
	}

	// Success to retrieve all information from one of the album hashes
	fmt.Printf("%s by %s: £%.2f [%d likes]\n", title, artist, price, likes)

}
