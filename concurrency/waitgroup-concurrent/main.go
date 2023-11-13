package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("this is test from goroutine 1")
		wg.Done()
	}()

	go func() {
		fmt.Println("this is test from goroutine 2")
		wg.Done()
	}()

	wg.Wait()
}
