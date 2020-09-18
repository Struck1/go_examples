package main

import (
	"fmt"
)

var deaultOffset = 10

func main() {

	defaultof := deaultOffset

	offset := 5

	val1, val2 := "emre", 8

	fmt.Println(offset, deaultOffset, val1, val2)

	defaultof = defaultof + deaultOffset

	val1, val2 = "emreeee", 88

	offset = 10

	fmt.Println(offset, defaultof,val1, val2)
}
