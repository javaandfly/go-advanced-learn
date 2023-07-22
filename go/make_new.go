package main

import "fmt"

type MyDemo struct {
	A string
}

func main() {
	myDemo := new(MyDemo)
	mapDemo := make(map[int]int)
	fmt.Println(mapDemo)
	fmt.Println(myDemo)
}
