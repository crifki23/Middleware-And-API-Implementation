package main

import (
	"chapter3-sesi2/handler"
	"errors"
	"unicode/utf8"
)

func IsValidName(name string) (string, error) {
	if utf8.RuneCountInString(name) < 5 {
		return "", errors.New("invalid name")
	}
	return "valid name", nil
}
func main() {
	handler.StartApp()
}
