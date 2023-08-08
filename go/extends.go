package main

import (
	"fmt"
	"unsafe"
)

type A struct {
}

type C struct {
	A
}

func main() {

	var c1 *C
	fmt.Printf("%p, size = %v\n", c1, unsafe.Sizeof(c1))
	c2 := new(C)
	fmt.Printf("%p, size = %v\n", c2, unsafe.Sizeof(c2))
	var c3 C
	fmt.Printf("%p, size = %v\n", &c3, unsafe.Sizeof(c3))
	c4 := C{}
	fmt.Printf("%p, size = %v\n", &c4, unsafe.Sizeof(c4))
	c5 := &C{}
	fmt.Printf("%p, size = %v\n", c5, unsafe.Sizeof(&c5))

}

// type B interface {
// 	Add()
// 	// A A
// }

// func add(b B) {
// 	fmt.Println("yes add")
// }
