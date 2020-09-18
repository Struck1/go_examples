package main

import (
	"fmt"
	"time"
)

func main() {

	go procces("older")
	procces("refund")

}

func procces(item string) {
	for i := 0; true; i++ {
		time.Sleep(time.Second)
		fmt.Println("Process", i, item)
	}
}
