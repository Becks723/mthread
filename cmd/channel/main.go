package main

import (
	"fmt"
	"time"
)

func main() {
	aniamlCh := make(chan string)

	go count(5, "sheep", aniamlCh)
	for msg := range aniamlCh {
		fmt.Println(msg)
	}
}

func count(n int, animal string, ch chan string) {
	for i := 0; i < n; i++ {
		ch <- animal
		time.Sleep(500 * time.Millisecond)
	}
	close(ch)
}
