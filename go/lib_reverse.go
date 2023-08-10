package main

import "fmt"

func main() {
	dog := Dog{name: "wangcai"}
	cat := Cat{name: "miaomiao"}
	speak(dog)
	speak(cat)
}

type Animal interface {
	Speak()
}

type Dog struct {
	name string
}

func (d Dog) Speak() {
	fmt.Printf("%s dog is wangwang\n", d.name)
}

type Cat struct {
	name string
}

func (c Cat) Speak() {
	fmt.Printf("%s cat is miaomiao", c.name)
}

func speak(a Animal) {
	a.Speak()
}
