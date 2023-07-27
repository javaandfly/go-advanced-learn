package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
		fmt.Println("First deferred statement")
	}()

	defer func() {

		// fmt.Println("Second deferred statement - will cause a panic")
		// 发生panic
		panic("Oops!")
	}()

	defer func() {
		fmt.Println("Third deferred statement")
	}()

}
