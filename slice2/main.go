package main

import (
	"fmt"
	"os"
)

var users = map[string]string{

	"305": "Sue",
	"204": "Bob",
	"631": "Jake",
	"073": "Tracy",
}

func getName(id string) (string, bool) {

	name, exist := users[id]
	return name, exist
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("User ıd was not passed")
		os.Exit(1)
	}
	name, exist := getName(os.Args[1])

	if !exist {
		fmt.Printf("error:  %v not found ", os.Args[1])
		os.Exit(1)
	}

	fmt.Println(name)
	fmt.Println(exist)
}
