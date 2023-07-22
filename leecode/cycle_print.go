package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	sign := 0
	a := 0
	wg.Add(2)
	go printA(&a, &sign)
	go printB(&a, &sign)
	wg.Wait()
}

func printA(a, sign *int) {
	for {
		if *a == 100 {
			wg.Done()
			return
		}
		if *sign == 0 {
			fmt.Printf("%d A", *a)
			fmt.Println()
			*a++
			*sign = 1
		}
	}

}

func printB(a, sign *int) {
	for {
		if *a == 100 {
			wg.Done()
			return
		}
		if *sign == 1 {
			fmt.Printf("%d B", *a)
			fmt.Println()
			*a++
			*sign = 0
		}
	}
}
