package model

import (
	"fmt"
)

type Post []struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (this Post) Get() string {
	var result string

	for _, post := range this {
		result += fmt.Sprintf("Id: %d\nTitle: %s\nBody: %s\n", post.Id, post.Title, post.Body)
	}

	return result
}
