package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// var waitgroup sync.WaitGroup

// func cleanup() {

// 	// r:=recover()
// 	// if r!= nil{
// 	// 	// do something
// 	// }
// 	defer waitgroup.Done()

// 	if r := recover(); r != nil {
// 		fmt.Println("recovered in cleanup", r)
// 	}
// }

// func say(s string) {
// 	defer cleanup()

// 	for i := 0; i < 3; i++ {
// 		fmt.Println(s)
// 		time.Sleep(time.Millisecond * 100)
// 		if i == 2 {
// 			panic("oh dear case 2")
// 		}
// 	}
// }

// func main() {
// 	waitgroup.Add(1)
// 	go say("hey")
// 	waitgroup.Add(1)
// 	go say("There")

// 	waitgroup.Wait()

// 	// waitgroup.Add(1)
// 	// go say("hi")
// 	// waitgroup.Wait()
// }
