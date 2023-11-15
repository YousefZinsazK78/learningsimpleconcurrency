package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	onceFunc := func() {
		fmt.Println("just once")
	}
	done := make(chan struct{})
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceFunc)
			done <- struct{}{}
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}
