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

	slice := []Salutaion{
		{"Ali", "Hi"},
		{"Jafar", "Ho"},
		{"Ghasem", "HU"},
	}

	for index, item := range slice {
		fmt.Println(index)
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
