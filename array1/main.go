package main

import (
	"fmt"
)

func arr1(arr [10]int) [10]int {

	for i := 0; i < len(arr); i++ {

		arr[i] = i + 1

	}

	return arr

}

func arr2(arr [10]int) [10]int {

	for i := 0; i < len(arr); i++ {

		arr[i] = arr[i] * arr[i]

	}

	return arr
}

func main() {

	var array [10]int
	firstArr := arr1(array)
	firstArr = arr2(firstArr)

	fmt.Println(firstArr)

}
