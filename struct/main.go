package main

import (
	"fmt"
)

type user struct {
	name    string
	age     int
	balance float64
	member  bool
}



func getUsers() []user {

	u1 := user{

		name:    "emre",
		age:     24,
		balance: 88.5,
		member:  true,
	}

	u2 := user{

		age:  22,
		name: "Nick",
	}

	u3 := user{

		"emree",
		25,
		0,
		false,
	}

	var u4 user

	u4.age = 28
	u4.name = "felan"

	return []user{u1, u2, u3, u4}

}

func main() {

	users := getUsers()

	fmt.Println(users)
	for i := 0; i < len(users); i++ {
		fmt.Printf("%v : %#v\n", i, users[i])
	}

}
