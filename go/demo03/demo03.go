package main

import (
	"fmt"
)

func main() {

	c := make(chan int, 2)
	c <- 1
	close(c)
	select {
	case a := <-c:
		fmt.Println(a)
		c <- 2
	}

}
