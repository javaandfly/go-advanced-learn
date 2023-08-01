package main

import (
	"fmt"
	"sync"
)

var wait sync.WaitGroup

func main() {
	wait.Add(2)
	go add(1)
	go add(1)
	wait.Wait()
}

func add(i int) {

	for {

		if i == 100 {
			wait.Done()
			return
		}
		fmt.Println(i)
		i++
	}
}
