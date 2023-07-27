package main

import (
	"fmt"
)

type A struct {
	C int
}

func main() {
	var a = "aaa"
	s := "aaa"
	fmt.Println(&a)
	fmt.Println(&s)
	// c := int64(1)
	c := a + s
	// var n int
	// n = 1
	fmt.Println(&c)

	// fmt.Println(&s == &c)
	// fmt.Println(unsafe.Sizeof(&s), unsafe.Sizeof(&c))
}
func B() *int {
	var a1 int
	return &a1
}
