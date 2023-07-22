package main

import "fmt"

type MyDemo struct {
	A string
}

var (
	size     = 1024
	max_size = size * 2
)

func main() {
	myDemo := new(MyDemo)
	mapDemo := make(map[int]int)
	fmt.Println(mapDemo)
	fmt.Println(myDemo)

	fmt.Println(size, max_size)

}
