package main

import "fmt"

type A struct {
	B
}

type C struct {
	B B
}

type B struct {
	a int
}

func main() {

	b := B{a: 1}
	a := A{B: b}
	c := C{B: b}

	fmt.Println(c.B.a)
	fmt.Println(a.a)

}
