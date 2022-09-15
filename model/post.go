package model

import (
	"fmt"
)

type Post []struct {
	UserId int    `json:userId`
	Id     int    `json:id`
	Title  string `json:title`
	Body   string `json:body`
}

func (posts Post) Get() string {
	result := ""

	for _, post := range posts {
		result += fmt.Sprintf("Id: %d\nTitle: %s\nBody: %s\n", post.Id, post.Title, post.Body)
	}

	return result
}
