package main

import (
	"fmt"
	"time"
)

func main() {

	intmap := make(map[int]int)
	for i := 0; i < 10000000; i++ {
		intmap[i] = i
	}
	start := time.Now()

	for j := 0; j < 10000000; j++ {
		_, ok := intmap[j]
		// fmt.Println(v)
		if !ok {
			fmt.Println("not find")
		}
	}
	duration := time.Since(start)
	fmt.Printf("代码执行时间：%s\n", duration)

}
