package main

import (
	"fmt"
	"http-api/client"
	"log"
)

func main() {
	post, err := client.Fetch()

	if err != nil {
		log.Println(err)
	}

	fmt.Println(post)
}
