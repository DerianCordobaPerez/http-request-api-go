package client

import (
	"encoding/json"
	"http-api/model"
	"io"
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

	bytes, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(bytes, &post)

	return post.Get(), err
}
