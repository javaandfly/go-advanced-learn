package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	demo := sync.Map{}

	go func() {
		for j := 0; j < 1000; j++ {
			demo.Store(j, j)
		}
	}()

	go func() {
		for j := 0; j < 1000; j++ {
			fmt.Println(demo.Load(j))
		}
	}()

	time.Sleep(time.Second * 1)
}
