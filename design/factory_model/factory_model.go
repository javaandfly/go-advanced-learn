package main

import "fmt"

type FactoryModel interface {
	Run()
}

type Ship struct {
	Name string `json:"name"`
}

func (s Ship) Run() {
	fmt.Println("Ship name: ", s.Name)
}

func main() {
	var ship Ship
	ship.Name = "123"
	add(ship)
}

func add(f FactoryModel) {
	f.Run()

}
