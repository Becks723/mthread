package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for {
			ch1 <- "sheep"
			time.Sleep(500 * time.Millisecond) // 每隔半秒
		}

	}()

	go func() {
		for {
			ch2 <- "cow"
			time.Sleep(2000 * time.Millisecond) // 每隔两秒
		}

	}()

	for {
		// select-case语句：
		// 进入select阻塞，哪个case通了进哪个case运行
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		}
	}
}
