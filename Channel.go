package main

import (
	"fmt"
)

func main() {
	fooVal := make(chan int)

	go foo(fooVal, 5)
	go foo(fooVal, 3)

	// v1 := <-fooVal
	// v2 := <-fooVal

	v1, v2 := <-fooVal, <-fooVal

	fmt.Println(v1)
	fmt.Println(v2)

}

func foo(c chan int, someValue int) {
	c <- someValue * 5
}
