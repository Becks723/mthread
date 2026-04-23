package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	query   string
	matches int

	workers       = 0
	maxWorkers    = 30
	workRequestCh = make(chan string)
	workDoneCh    = make(chan struct{})
	foundMatchCh  = make(chan struct{})
)

func main() {
	query = "test"
	start := time.Now()

	workers = 1
	go search("/home/becks723", true)
	waitForWorkers()
	fmt.Println(strconv.Itoa(matches) + " matches")
	fmt.Println(time.Since(start))
}

func search(path string, master bool) {
	files, err := os.ReadDir(path)
	if err == nil {
		for _, file := range files {
			if file.IsDir() {
				if workers < maxWorkers {
					workRequestCh <- path + "/" + file.Name()
				} else {
					search(path+"/"+file.Name(), false)
				}
			} else if file.Name() == query {
				foundMatchCh <- struct{}{}
			}
		}
	}
	if master {
		workDoneCh <- struct{}{}
	}
}

func waitForWorkers() {
	for {
		select {
		case path := <-workRequestCh:
			workers++
			go search(path, true)
		case <-workDoneCh:
			workers--
			if workers == 0 {
				return
			}
		case <-foundMatchCh:
			matches++
		}
	}
}
