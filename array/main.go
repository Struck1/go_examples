package main

import (
	"fmt"
)

func createArray() [10]int {

	var arr [10]int
	return arr

}

func compArray() (bool, bool, bool) {

	var arr1 [5]int
	fmt.Printf("%#v \n", arr1)
	arr2 := [5]int{0}
	fmt.Printf("%#v \n", arr2)
	arr3 := [...]int{0, 0, 0, 0, 0}
	arr4 := [5]int{0, 0, 0, 0, 9}

	return arr1 == arr2, arr1 == arr3, arr1 == arr4
}

func compArrr1() (bool, bool, [10]int) {

	var arr11 [10]int
	arr22 := [...]int{9: 0}
	arr33 := [10]int{1, 4: 5, 9: 10}

	return arr11 == arr22, arr11 == arr33, arr33
}

func indexArray() string {

	arr := [...]string{
		"a",
		"c",
		"b",
	}

	return fmt.Sprintln(arr[0], arr[2], arr[1])

}

func message() string {

	arr := [4]int{1, 2, 3, 4}
	m := ""

	for i := 0; i < len(arr); i++ {

		arr[i] = arr[i] * arr[i]

		m += fmt.Sprintf("%v : %v \n", i, arr[i])

	}

	return m

}

func main() {

	fmt.Printf("%#v \n", createArray())

	comp1, comp2, comp3 := compArray()

	fmt.Println("[5]int == [5]int{0}       :", comp1)
	fmt.Println("[5]int == [...]int{0, 0, 0, 0, 0}:", comp2)
	fmt.Println("[5]int == [5]int{0, 0, 0, 0, 9} :", comp3)

	comp11, comp22, arr33 := compArrr1()
	fmt.Println("[10]int == [...]{9:0}       :", comp11)
	fmt.Println("[10]int == [10]int{1, 9: 10, 4: 5}}:", comp22)
	fmt.Println("arr3               :", arr33)

	fmt.Print(indexArray())

	fmt.Print(message())

}
