package main

import (
	"fmt"
)

type A struct {
	C int
}

func main() {
	// var a = "aaa"
	// s := "aaa"
	// c := int64(1)

	var n int
	n = 1
	fmt.Println(n)

	// fmt.Println(&s == &c)
	// fmt.Println(unsafe.Sizeof(&s), unsafe.Sizeof(&c))
}
func B() *int {
	var a1 int
	return &a1
}
