package main

import (
	"fmt"
)

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human
	specialty string
}

func birEkle(a *int) int {

	*a = *a + 1
	return *a
}

func GoodBye() {

	fmt.Println("Hello")
}

func main() {

	emre := Student{Human{"EMRE", 25, 100}, "Mek"}

	x := 3

	fmt.Println("x = ", x)

	x1 := birEkle(&x)

	fmt.Println("X + 1 :", x1)
	fmt.Println("X", x)

	defer GoodBye() // main içinde tüm işlemler bittikden sonra defer tanımlı fonk çalıştı defer olmasa hello - emre olacaktı
	fmt.Println("Emre")

	fmt.Println(emre.age)
	fmt.Println(emre.name)
	fmt.Println(emre.weight)
	fmt.Println(emre.specialty)

}
