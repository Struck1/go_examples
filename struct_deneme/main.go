package main

import "fmt"

type Skills []string

type Human struct {
	name   string
	age    int
	height float64
	phone  string // human bir tele sahip
}

type Student struct {
	Human
	Skills
	int
	specialty string
	phone     string // student de bir tele sahip
}

func main() {

	emre := Student{Human: Human{"EMRE", 24, 175.5, "555-444-333"}, specialty: "FALAN", phone: "555-555-555"}

	fmt.Println(emre.age)
	fmt.Println(emre.name)
	fmt.Println(emre.height)
	fmt.Println(emre.specialty)
	fmt.Println(emre.Human.phone)
	fmt.Println(emre.phone)

	emre.Skills = []string{"felan"}
	fmt.Println(emre.Skills)
	emre.Skills = append(emre.Skills, "felan1", "felan2")
	fmt.Println(emre.Skills)
	emre.int = 3
	fmt.Println(emre.int)
}
