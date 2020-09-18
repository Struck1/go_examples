package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type API struct {
	Message string `json:"message"`
}

func UrlParam(w http.ResponseWriter, r *http.Request) {

	urlParams := mux.Vars(r)

	id := urlParams["id"]

	message := "Kullanıcı adı" + id

	mess := API{message}

	output, err := json.Marshal(mess)

	if err != nil {
		fmt.Println("err")
	}

	fmt.Fprint(w, string(output))
}

func main() {

	route := mux.NewRouter()

	route.HandleFunc("/api/{id:[0-9]+}", UrlParam)

	http.Handle("/", route)

	http.ListenAndServe(":9090", nil)
}
