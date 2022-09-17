package main

import (
	"fmt"
	"github.com/abraaolincoln/router"
	"io"
	"net/http"
)

func main() {
	mux := router.NewRestMux()
	mux.Get("/", echo)
	mux.Get("/:id", echoPathVariable)
	http.ListenAndServe(":3333", mux)
}

func echo(w http.ResponseWriter, r *http.Request, extraInfo *router.ExtraInfo) {
	io.WriteString(w, "Hello")
}

func echoPathVariable(w http.ResponseWriter, r *http.Request, extraInfo *router.ExtraInfo) {
	fmt.Println(extraInfo.PathVariables[":id"])
	io.WriteString(w, "Hello")
}
