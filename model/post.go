package model

import (
	"fmt"
)

type Post []struct {
	UserId string `json:userId`
	Id     string `json:id`
	Title  string `json:title`
	Body   string `json:body`
}

func (post Post) TextOutput() string {
	return fmt.Sprintf("Id: %s\nTitle: %s\nBody: %s\n", post[0].Id, post[0].Title, post[0].Body)
}
