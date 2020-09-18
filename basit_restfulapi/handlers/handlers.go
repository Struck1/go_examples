package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	. "../dataloaders"
	. "../models"
)

func Run() {

	http.HandleFunc("/", Handler)

	http.ListenAndServe(":9090", nil)
}

func Handler(res http.ResponseWriter, r *http.Request) {

	page := Page{ID: 7, Name: "Kullanıcılar", Description: "Kullanıcı Listesi", URL: "/users"}

	users := LoadUsers()

	interest := LoadInterests()

	interestMappings := LoadInterestMapping()

	var newUsers []User

	fmt.Println(users)

	for _, user := range users {
		fmt.Println(user)
		for _, mapping := range interestMappings {
			if user.ID == mapping.UserID {
				for _, interstt := range interest {
					if mapping.InterestID == interstt.ID {
						user.Interests = append(user.Interests, interstt)
					}

				}
			}
		}

		newUsers = append(newUsers, user)
	}

	viewModel := UserViewModels{Page: page, Users: newUsers}

	data, _ := json.Marshal(viewModel)

	res.Write([]byte(data))

}
