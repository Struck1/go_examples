package main

import (
	"fmt"
)

type Account struct {
	AccoundNumber int
	FirstName     string
	LastName      string
}

func (account *Account) GetFullName() string {

	fmt.Println(*account)

	return account.FirstName + " " + account.LastName

}

func (account *Account) ChangeName(name string) { // pointer 

	account.FirstName = name

}

func main() {

	account := Account{22, "emre", "biçen"}
	account.ChangeName("falan")
	fmt.Println(account.GetFullName())

}
