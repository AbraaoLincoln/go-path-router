package main

import (
	"fmt"
	"github.com/abraaolincoln/router"
	"net/http"
)

func main() {
	path := "/user/123/posts"
	handle := func(w http.ResponseWriter, r *http.Request) { fmt.Println("handler1") }
	root := router.NewNode()
	tree := router.Tree{Root: &root}

	tree.Insert(path, handle)

	path = "/users/profile"
	handle = func(w http.ResponseWriter, r *http.Request) { fmt.Println("handler2") }

	tree.Insert(path, handle)
}
