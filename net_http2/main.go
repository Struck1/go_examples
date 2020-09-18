package main

import (
	"fmt"
	"log"
	"net/http"
)

type messageHandler struct {
	message string
}

func (x *messageHandler) ServeHTTP(res http.ResponseWriter, r *http.Request) {

	fmt.Fprint(res, x.message)
}

func (x *messageHandler) ChangeName(name string) {

	x.message = name
}

func main() {

	message := &messageHandler{"emre"}
	message.ChangeName("emre değil")
	
	message2 := &messageHandler{"bicen"}

	mux := http.NewServeMux()

	mux.Handle("/mes1", message)
	mux.Handle("/mes2", message2)

	log.Println("Listening...")

	http.ListenAndServe(":9090", mux)

}
