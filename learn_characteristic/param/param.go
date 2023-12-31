package main

type B struct {
}

type C struct {
	B B
}

func A(b *B) *C {

}
