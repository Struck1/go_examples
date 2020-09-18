package main

import "fmt"

func sum(a []int, c chan int) {
	total := 0

	for _, v := range a {
		total += v
	}

	c <- total
}

func main() {

	a := []int{7, 2, 6, 8, 9, 4, -4, 7}

	c := make(chan int)

	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)

	//Buffered channels
	d := make(chan int, 2) // change 2 to 1 will have runtime error, but 3 is fine
	d <- 1
	d <- 2
	//d <- 3
	fmt.Println(<-d)
	fmt.Println(<-d)
	//fmt.Println(<-d)

}
