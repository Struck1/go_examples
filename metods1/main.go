package main

import (
	"fmt"
)

type Numbers struct {
	X, Y int
}

func (numbers Numbers) Sum() int {

	sum := numbers.Y + numbers.X

	return sum
}

func main() {

	number := Numbers{5, 5}

    fmt.Println(number.Sum())

}
