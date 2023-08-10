package main

import (
	"fmt"
	"sync"
)

func myFunc() {
	fmt.Println("hello world")
}

func main() {
	var once sync.Once
	once.Do(myFunc)
}
