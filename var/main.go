package main

import (
	"fmt"
	"time"
)

var foo string = "bar" //global

// declare multible variables
var (
	ddebug bool = false

	level string = "info"

	startTime time.Time = time.Now()
)

// Skipping the Type or Value When Declaring Variables
var (
	debugg bool

	levell = "ingo"

	startTimee = time.Now()
)

func getName() (string, string, time.Time) {

	return "emre", "bicen", time.Now()
}

func main() {

	var name string = "emre"

	isim := "emre" // short	 variable declaration

	falan := "filan"

	variable1, variable2, variable3 := 1, 2, 3 // multible short variable declaration

	name1, name2, startt := getName()

	//fmt.Println(name, foo, ddebug, level, startTime)
	fmt.Println(name, debugg, levell, startTimee, isim, falan, variable1, variable2, variable3, name1, name2, startt)
}
