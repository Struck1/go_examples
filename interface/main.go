package main

import "fmt"

type IAccount interface {
	GetFullName() string
	GetAccount() int
}

type Account struct {
	FirstName     string
	LastName      string
	AccountNumber int
}

func (account Account) GetFullName() string {

	return account.FirstName + " " + account.LastName
}

func (account Account) GetAccountNumber() int { // hata !!!!

	return account.AccountNumber
}

func (account Account) GetAccount() int {

	return account.AccountNumber
}

func printFullName(account IAccount) {

	fmt.Println(account.GetFullName())
}

func printScreen(value interface{}) {
	hello, ok := value.(int) // string

	if ok {
		fmt.Println(hello)
	} else {
		fmt.Println("Tip dönüşümü olmadı")
	}
}

func main() {

	var myAccount IAccount = Account{"EMRE", "BİÇEN", 222} // BU şekilde de tanımlaya biliriz interface
	account := Account{"emre", "biçen", 22}

	printFullName(account)
	fmt.Println(myAccount.GetFullName())

	var name = "Emre"

	printScreen(name)

}
