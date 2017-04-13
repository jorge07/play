package post

import (
	"reflect"
	"encoding/json"
	"log"
	"fmt"
)

type Post struct {
	UserId int `json:"userId"`
	Id     int `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (p Post) IsEmpty() bool {

	return reflect.DeepEqual(Post{}, p)
}

func (p Post) Render() {

	s, err := json.MarshalIndent(p, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", s)
	fmt.Println("")
}
