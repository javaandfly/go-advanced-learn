package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	chan1 := make(chan bool, 1)
	chan2 := make(chan bool, 1)
	chan2 <- false
	v := 1
	wg.Add(2)
	go PtfInt(v, chan1, chan2)
	char := 'A'
	go PtfChar(char, chan1, chan2)
	wg.Wait()

}

func PtfInt(v int, chan1, chan2 chan bool) {
	for {
		<-chan2
		if v == 28 {
			wg.Done()

			return
		}
		for i := 0; i < 2; i++ {
			fmt.Print(v)
			v++
		}
		chan1 <- false
	}

}

func PtfChar(char rune, chan1, chan2 chan bool) {
	for {
		<-chan1
		if char == 'Z' {
			wg.Done()
			return
		}
		for i := 0; i < 2; i++ {
			fmt.Print(char)
			char++
		}
		chan2 <- false

	}

}
