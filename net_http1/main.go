package main

import (
	"io"
	"log"
	"net/http"
)

type falan int

func (x falan) ServeHTTP(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "falan")
}

type filan int

func (x filan) ServeHTTP(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "filan")
}

func main() {

	var i falan
	var j filan

	mux := http.NewServeMux()
	mux.Handle("/ilk", i)
	mux.Handle("/son", j)

	err := http.ListenAndServe(":9090", mux)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
