package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("this is just a test")
	go func() {
		fmt.Println("this is just a test from first goroutine")
		wg.Done()
	}()
	go func() {
		fmt.Println("this is just a test from second goroutine")
		wg.Done()
	}()
	wg.Wait()
}
