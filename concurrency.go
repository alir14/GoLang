package main

import (
	"fmt"
	"sync"
	"time"
)

var waitgroup sync.WaitGroup

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}

	waitgroup.Done()
}

func main() {
	waitgroup.Add(1)
	go say("hey")
	waitgroup.Add(1)
	go say("There")
	// waitgroup.Wait()

	waitgroup.Add(1)
	go say("hi")
	waitgroup.Wait()
}
