package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/foo", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "hi foo")
	}).Methods("GET")
	r.HandleFunc("/users/{user:[a-z]+}", func(w http.ResponseWriter, req *http.Request) {
		user := mux.Vars(req)["user"]
		fmt.Fprintf(w, "hi %s", user)
	}).Methods("GET")
	http.ListenAndServe(":8000", r)
}
