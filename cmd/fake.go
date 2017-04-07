package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"net/http"
	"strings"
	"encoding/json"
	"time"
	"log"
	"io/ioutil"
	"reflect"
)

type Post struct {
	UserId int `json:"userId"`
	Id int `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
}

func (p Post ) IsEmpty() bool {

	return reflect.DeepEqual(Post{}, p)
}

func (p Post) render() {

	s, err := json.MarshalIndent(p, "", "    ")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", s)
	fmt.Println("")
}

func hydrate(res *http.Response, p *Post)  {
	body, readErr := ioutil.ReadAll(res.Body)

	if readErr != nil {
	log.Fatal(readErr)
	}

	jsonErr := json.Unmarshal(body, &p)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	if p.IsEmpty() == true {
		log.Fatal("Empty response")
	}
}

func call(url string) *http.Response {

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	return res
}

var fakeCmd = &cobra.Command{
	Use:   "fake",
	Short: "Display fake data given an identifer",
	Long: "Display fake data from https://jsonplaceholder.typicode.com/posts fir the id given",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Print("Page its necesary");
		}

		url := "https://jsonplaceholder.typicode.com/posts/" + strings.Join(args, "")

		var res = call(url);

		var p = &Post{}
		hydrate(res, p)

		p.render()

	},
}

func init() {
	RootCmd.AddCommand(fakeCmd)
}
