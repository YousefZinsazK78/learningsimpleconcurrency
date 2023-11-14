package main

import (
	"fmt"
	"sync"
)

var counter int

func main() {
	m := map[int]int{}

	var mu = new(sync.RWMutex)

	go writeloop(m, mu)
	go readloop(m, mu)
	go readloop(m, mu)
	go readloop(m, mu)
	go readloop(m, mu)
	go readloop(m, mu)
	go readloop(m, mu)

	block := make(chan struct{})

	<-block
}

func readloop(m map[int]int, mu *sync.RWMutex) {
	for {
		mu.RLock()
		for k, v := range m {
			fmt.Println(k, "-", v)
		}
		mu.RUnlock()
	}
}

func writeloop(m map[int]int, mu *sync.RWMutex) {
	for {
		for i := 0; i < 100; i++ {
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}
	}
}
