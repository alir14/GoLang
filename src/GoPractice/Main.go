package main

import (
	"fmt"
)

const (
	PI = 3.14
)

type Salutaion struct {
	name     string
	grreting string
}

type MyPrinter func(string)

func main() {

	// var prefixMap map[string]string
	// prefixMap = make(map[string]string)
	// OR
	// prefixMap := make(map[string]string)

	// prefixMap["Ali"] = "Agha"
	// prefixMap["Jafar"] = "Doctor"
	// prefixMap["Ghasem"] = "Mohandes"

	// OR

	prefixMap := map[string]string{
		"Ali":     "Agha",
		"Jafar":   "Doctor",
		"Ghasem":  "Mohandes",
		"Nazanin": "Khanom",
	}

	delete(prefixMap, "Nazanin")

	slice := []Salutaion{
		{"Ali", "Hi"},
		{"Jafar", "Ho"},
		{"Ghasem", "HU"},
		{"Nazanin", "He"},
	}

	for index, item := range slice {

		if value, exists := prefixMap[item.name]; exists {
			item.grreting = value
		} else {
			item.grreting = "Dude"
		}

		fmt.Print(index, ") ")
		item.grreting = prefixMap[item.name]
		Greeting(item, CreatePrintFunction("!!!"))
	}

	// var salutation = Salutaion{"Ali", "Hello"}
	// Greeting(salutation, CreatePrintFunction("!!!"))
}

func Greeting(Salutaion Salutaion, do MyPrinter) {
	message, alternative := CreateMessage(Salutaion.name, Salutaion.grreting, "YO")
	do(message)
	do(alternative)
}

// func Print(input string) {
// 	fmt.Print(input)
// }

func CreatePrintFunction(custom string) MyPrinter {
	return func(input string) {
		fmt.Println(input + custom)
	}
}

// func PrintCustom(input string, custom string) {
// 	fmt.Println(input + custom)
// }

func CreateMessage(name string, greeting ...string) (message string, alternative string) {
	message = greeting[0] + " " + name
	alternative = "Hey! " + name
	return
}
