package main

import (
	"fmt"
)

type Dog struct {
	Name  string
	Color string
}

//Method for object specify the object the refer to and then
//the method name and rest of
func (d Dog) Call() {
	fmt.Printf("Come here %s dog, %s \n", d.Color, d.Name)
}

func main() {
	Spot := Dog{Name: "Spot", Color: "brown"}

	Spot.Call()
}
