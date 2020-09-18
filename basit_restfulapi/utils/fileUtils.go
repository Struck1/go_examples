package utils

import (
	"errors"
	"io/ioutil"
	"log"
)

func ReadFile(filename string) (string, error) {

	if IsEmpty(filename) {
		return "", errors.New("Dosya adı kullanılamaz")
	}

	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	return string(bytes), nil
}
