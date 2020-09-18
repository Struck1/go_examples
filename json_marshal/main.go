package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type API struct {
	Message string `json:"message"`
}

type Person struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
}

func main() {

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {

		message := API{"HELLO"}

		output, _ := json.Marshal(message)

		fmt.Fprint(w, string(output))

	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {

		users := []Person{

			Person{ID: 1, FirstName: "emre", LastName: "biçen", Age: 24},
			Person{ID: 1, FirstName: "emre", LastName: "biçen", Age: 24},
			Person{ID: 1, FirstName: "emre", LastName: "biçen", Age: 24},
		}

		message := users

		output, _ := json.Marshal(message)

		fmt.Fprint(w, string(output))

	})

	http.ListenAndServe(":9090", nil)

}
