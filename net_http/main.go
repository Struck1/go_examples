package main

import (
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Index page"))
	// w.WriteHeader(http.StatusOK)

	u := r.URL.Path

	w.Write([]byte(u))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("About page"))
}

func main() {

	http.HandleFunc("/", indexHandler)

	http.HandleFunc("/index", indexHandler)

	http.HandleFunc("/about", aboutHandler)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
