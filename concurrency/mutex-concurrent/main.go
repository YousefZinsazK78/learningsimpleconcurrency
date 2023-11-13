package main

import (
	"fmt"
	"time"
)

func main() {
	sum := 0

	inc := func() {
		sum += 10
	}

	dec := func() {
		sum -= 10
	}

	for i := 0; i < 5; i++ {
		go inc()
	}

	go dec()

	time.Sleep(time.Second)
	fmt.Println(sum)
}
