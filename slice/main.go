package main

import (
	"fmt"
)

func main() {

	//carArray := [3]string{"bmw", "ferrari", "opel"}

	carSlice := []string{"bmw", "ferrari", "opel"} // boyut belli değil

	for index, carBrand := range carSlice {

		fmt.Println(index, carBrand)
	}

	teams := []string{}
	fmt.Printf("uzunluk: %d, kapasite: %d \n ", len(teams), cap(teams))

	teams = append(teams, "emre")
	fmt.Printf("uzunluk: %d, kapasite: %d \n ", len(teams), cap(teams))

	teams = append(teams, "biçen")
	fmt.Printf("uzunluk: %d, kapasite: %d \n ", len(teams), cap(teams))

	teams = append(teams, "hi")
	fmt.Printf("uzunluk: %d, kapasite: %d \n ", len(teams), cap(teams))

	fmt.Println(teams)

}
