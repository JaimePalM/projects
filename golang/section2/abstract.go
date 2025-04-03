package main

import (
	"fmt"
)

type Car interface {
	Drive()
	Stop()
}

type Lamborghini struct {
	LamborghiniModel string
}

// This is a constructor, we return a Lamborghini as a Car, so if Lamborghini does not implement the Car interface, we will get a compile error
func NewModel(arg string) Car {
	return &Lamborghini{arg}
}

type Chevy struct {
	ChevyModel string
}

// sintax: func (parameter) method() {}
// In this case the parameter is a pointer to the struct in order to be able to access its fields
func (l *Lamborghini) Drive() {
	fmt.Println("Lamborghini", l.LamborghiniModel, "is driving")
}

func (l *Lamborghini) Stop() {
	fmt.Println("Lamborghini", l.LamborghiniModel, "is stopping")
}

func (c *Chevy) Drive() {
	fmt.Println("Chevy is driving")
	fmt.Println(c.ChevyModel)
}

// EMPTY INTERFACES

func Anything(anything interface{}) {
	fmt.Println(anything)
}

func main() {

	l := new(Lamborghini)
	l.LamborghiniModel = "Aventador"
	l.Drive()

	c := new(Chevy)
	c.ChevyModel = "Camaro"
	c.Drive()

	l2 := NewModel("Gallardo")
	l2.Drive()
	l2.Stop()

	// Empty Interfaces
	Anything(43.2)
	Anything("Hello")
	Anything(struct{}{})

	mymap := make(map[string]interface{})

	mymap["name"] = "John"
	mymap["age"] = 30

	fmt.Println(mymap)

}
