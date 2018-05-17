package main

import (
	"fmt"
	"net/http"
)

const x int = 5

func main() {
	// fmt.Println("Hello Go language")
	// foo()
	// fmt.Println("Random Number ", rand.Intn(150))
	// var num1 float64 = 2.4
	// var num2 float64 = 8.43
	// var num1, num2 float64 = 5.6, 9.4
	// num1, num2 := 5.6, 9.6
	// fmt.Println(add(num1, num2))

	// x := 15
	// a := &x
	// fmt.Println(a)
	// fmt.Println(*a)

	http.HandleFunc("/", index_Handler)
	http.HandleFunc("/about", about_Handler)
	http.ListenAndServe(":8000", nil)
}

func index_Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Uhm..., what do you want ?")
}

func about_Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Uhm..., is it what you are looking for ?")
}

// func foo() {
// 	fmt.Println("The square root of 4 is ", math.Sqrt(4))
// }
// func add(x float64, y float64) float64 {
// 	return x + y
// }
// func multiple(a, b string) (string, string) {
// 	return a, b
// }
