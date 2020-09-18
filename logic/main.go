package main

import (
	"fmt"
)

func main() {

	var total float64 = 2 * 13
	fmt.Println("Sub :", total)

	total = total + (4 * 2.25)
	fmt.Println("Sub :", total)

	total = total - 5
	fmt.Println("Sub :", total)

	tip := total * 0.1
	fmt.Println("Tip :", tip)

	total = total + tip
	fmt.Println("Total :", total)

	split := total / 2
	fmt.Println("Split : ", split)

	visitCount := 24

	visitCount = visitCount + 1

	remainder := visitCount % 5

	if remainder == 0 {

		fmt.Println("With this visit, you've earned a reward")
	}

	count := 10

	count += 10

	count--

	count -= 5

	fmt.Println("Count :", count)

	visits := 15

	fmt.Println("Firt visit", visits == 1)
	fmt.Println("Firt visit", visits != 1)
	fmt.Println("Firt visit", visits >= 15 && visits < 16)
	fmt.Println("Firt visit", visits >= 15 || visits < 11)

	var emails []string
	fmt.Printf("Emails : %#v \n", emails)

	var message string
	fmt.Printf("Message : %#v \n", message)

}
