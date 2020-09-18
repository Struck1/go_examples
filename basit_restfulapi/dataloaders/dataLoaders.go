package dataloaders

import (
	"encoding/json"
	. "../models"
	util "../utils"
)

func LoadUsers() []User {

	string_data, _ := util.ReadFile("../json/users.json")

	var data []User

	json.Unmarshal([]byte(string_data), &data)

	return data

}

func LoadInterests() []Interest {

	string_data, _ := util.ReadFile("../json/interests.json")

	var data []Interest

	json.Unmarshal([]byte(string_data), &data)

	return data

}

func LoadInterestMapping() []InterestMapping {

	string_data, _ := util.ReadFile("../json/userInterestMappings.json")

	var data []InterestMapping

	json.Unmarshal([]byte(string_data), &data)

	return data
}
