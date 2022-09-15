package client

import (
	"encoding/json"
	"http-api/model"
	"log"
	"net/http"
)

func Fetch() (string, error) {
	url := "https://jsonplaceholder.typicode.com/posts"

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal("Error while get data", err)
	}

	defer resp.Body.Close()

	var post model.Post

	if err := json.NewDecoder(resp.Body).Decode(&post); err != nil {
		log.Fatal("Error while generate json ", err)
	}

	return post.Get(), nil
}
