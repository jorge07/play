package post

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"time"
)

type PostRepository struct {
	url string
	client http.Client
}


func (r *PostRepository) Call(identifier string) *Post {

	req, err := http.NewRequest(http.MethodGet, r.url + identifier, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := r.client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	return r.hydrate(res)
}

func  (r *PostRepository) hydrate(res *http.Response) *Post {

	p := &Post{}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	if jsonErr := json.Unmarshal(body, &p); jsonErr != nil {
		log.Fatal(jsonErr)
	}

	if p.IsEmpty() == true {
		log.Fatal("Empty response")
	}

	return p
}


func GetRepository() *PostRepository  {

	r := &PostRepository{}
	r.url = "https://jsonplaceholder.typicode.com/posts/"
	r.client = http.Client{
		Timeout: time.Second * 2,
	}

	return r
}