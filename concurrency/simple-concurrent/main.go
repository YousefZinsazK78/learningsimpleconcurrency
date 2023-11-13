package main

import (
	"fmt"
	"time"
)

// simple goroutine and concurrency in golang
func main() {

	go RunForLoop()
	time.Sleep(time.Second * 1)

}

func RunForLoop() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

}
