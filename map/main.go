package main

import (
	"fmt"
	"os"
)

var values = map[string]string{

	"100": "emre",
	"101": "biçen",
	"102": "felan",
}

func getUsers() map[string]string {

	values["103"] = "heh"

	return values

}

func deleteUsers(id string) {
	delete(values, id)
}

func slicee() []string {

	fist := []string{"Goood", "Good", "Bad", "Good", "Good"}

	second := append(fist[:2], fist[3:]...)

	return second

}

func main() {

	userıd := os.Args[1]
	deleteUsers(userıd)

	fmt.Println("Users: ", getUsers())
	fmt.Println(slicee())


}
