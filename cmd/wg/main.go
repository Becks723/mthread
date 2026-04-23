package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go count(5, "cow")
	go count(5, "sheep")

	wg.Wait()
}

func count(n int, animal string) {
	defer wg.Done()

	for i := 0; i < n; i++ {
		fmt.Println(animal)
		time.Sleep(500 * time.Millisecond)
	}

}
