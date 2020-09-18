package main

import (
	"fmt"
	"os"
)



func getpassed() []string {

	var args []string

	for i := 0; i < len(os.Args); i++ {

		args = append(args, os.Args[i])

	}

	return args

}

func getLocal(extralocal []string) []string {

	var locales []string

	locales = append(locales, "en_US", "fr_FR")

	locales = append(locales, extralocal...)

	return locales

}

func main() {

	locales := getLocal(getpassed())
	fmt.Println("Locales to use:", locales)

}
