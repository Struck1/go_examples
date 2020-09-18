package main

import (
	"fmt"
)

func main() {

	var count *int

	count2 := new(int)

	countTemp := 5

	count3 := &countTemp

	fmt.Printf("count1: %#v\n", count)
	fmt.Printf("count2: %#v\n", count2)
	fmt.Printf("count3: %#v\n", count3)
	fmt.Printf("count3: %#v\n", *count3)

}
