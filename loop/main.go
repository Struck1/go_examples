package main

import (
	"errors"
	"fmt"
	"strconv"
)

func validate(input int) error {

	if input < 0 {
		return errors.New("input can't be negative number")
	} else if input > 100 {
		return errors.New("input can't be over 100")
	} else {
		return nil
	}
}

func main() {

	variable := 101

	if err := validate(variable); err != nil {
		fmt.Println(err)
	} else if variable%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	sum := 1

	for sum <= 100 {

		if sum%5 == 0 && sum%3 == 0 {

			res2 := strconv.Itoa(sum)
			res2 = "FizzBuzz"
			fmt.Println(res2)

		} else if sum%5 == 0 {

			res1 := strconv.Itoa(sum)
			res1 = "Buzz"
			fmt.Println(res1)

		} else if sum%3 == 0 {
			res := strconv.Itoa(sum)
			res = "Fizz"
			fmt.Println(res)

		} else {
			fmt.Println(sum)
		}

		sum++

	}

	arr := []string{"Jim", "Jane", "Joe", "June"} // arr name - arr len - type

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	config := map[string]string{ // map name - key type - value type
		"debug":     "1",
		"log level": "warn",
		"version":   "1.2",
	}

	for key, value := range config { // The range keyword is mainly used in for loops in order to iterate over all the elements of a map, slice, channel, or an array

		fmt.Println(key, "=", value)
	}

	words := map[string]int{
		"Gonna": 3,
		"You":   3,
		"Give":  2,
		"Never": 1,
		"Up":    4,
	}

	max := 0
	name := ""

	for key, value := range words {

		if value > max {
			max = value
			name = key
		}
	}
	fmt.Println("Most popular word :", name)
	fmt.Println("With a count of :", max)
}
