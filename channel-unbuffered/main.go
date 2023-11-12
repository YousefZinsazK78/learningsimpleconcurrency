package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan string)
	go writer(ch, &wg)
	go reader(ch, &wg)
	wg.Wait()
}

func writer(ch chan<- string, wg *sync.WaitGroup) {
	ch <- "yousef"
	wg.Done()
}

func reader(ch <-chan string, wg *sync.WaitGroup) {
	fmt.Println(<-ch)
	wg.Done()
}
