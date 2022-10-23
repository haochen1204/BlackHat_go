package main

import (
	"fmt"
	"net/http"
)

type router struct {
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/a":
		fmt.Fprint(w, "executing /a")
	case "/b":
		fmt.Fprint(w, "executing /b")
	case "/c":
		fmt.Fprint(w, "executing /c")
	default:
		http.Error(w, "404 Not Found", 404)
	}
}

func main() {
	var r router
	http.ListenAndServe(":8000", &r)
}
