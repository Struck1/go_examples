package main

import (
	"fmt"
)

func getSlice() ([]int, []int, []int) {

	var s1 []int

	s2 := make([]int, 10)

	s3 := make([]int, 10, 50)

	return s1, s2, s3

}

func linked() (int, int, int) {

	s1 := []int{1, 2, 3, 4, 5}

	s2 := s1

	s3 := s1[:]

	s1[3] = 99

	return s1[3], s2[3], s3[3]

}

func noLink() ([]int, []int) {

	s1 := []int{1, 2, 3, 4, 5}

	s2 := s1

	s1 = append(s1, 6)

	s1[3] = 99

	return s1, s2

}

func capLinked() ([]int, []int) {

	s1 := make([]int, 5, 10)

	s1[0], s1[1], s1[2], s1[3], s1[4] = 1, 2, 3, 4, 5

	s2 := s1

	s1 = append(s1, 6)

	s1[3] = 99

	return s1, s2
}

func capNoLink() ([]int, []int) {

	s1 := make([]int, 5, 10)

	s1[0], s1[1], s1[2], s1[3], s1[4] = 1, 2, 3, 4, 5

	s2 := s1

	s1 = append(s1, []int{10: 11}...)

	s1[3] = 99

	return s1, s2

}

func copyNoLinke() (int, []int, int) {

	s1 := []int{1, 2, 3, 4, 5}

	s2 := make([]int, len(s1))

	fmt.Println(s2)

	copied := copy(s2, s1)

	s1[3] = 99

	return s1[3], s2, copied

}

func appendNoLink() ([]int, []int) {

	s1 := []int{1, 2, 3, 4, 5}

	s2 := append([]int{}, s1...)

	s1[3] = 99

	return s1, s2
}

func main() {

	s1, s2, s3 := getSlice()

	fmt.Printf("S1 : len = %v cap = %v\n", len(s1), cap(s1))
	fmt.Printf("S2 : len = %v cap = %v\n", len(s2), cap(s2))
	fmt.Printf("S3 : len = %v cap = %v\n", len(s3), cap(s3))

	fmt.Println(linked())
	fmt.Println(noLink())
	fmt.Println(capLinked())
	fmt.Println(capNoLink())
	fmt.Println(copyNoLinke())
	fmt.Println(appendNoLink())

}
