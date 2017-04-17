package cmd

import (
	"testing"
	"github.com/jorge07/play/post"
	"log"
	"fmt"
)

func TestFake(t *testing.T) {
	repo := post.GetRepository()

	p := repo.Call("1");

	if (p.Id != 1) {
		log.Fatal(fmt.Sprintf("Post.Id must be 1 and its %d", p.Id))
	}
}
