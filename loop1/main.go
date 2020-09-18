package main

import (
	"fmt"
)

func main() {

	nums := []int{5, 8, 2, 4, 0, 1, 3, 7, 9, 6}

	//fmt.Println(reflect.TypeOf(nums))

	fmt.Println("Before:", nums)

	for swapp := true; swapp; {

		swapp = false

		for i := 1; i < len(nums); i++ {

			if nums[i] < nums[i-1] {

				nums[i], nums[i-1] = nums[i-1], nums[i]

				swapp = true
			}

		}
	}

	fmt.Println("After: ", nums)

}
