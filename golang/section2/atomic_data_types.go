package main

import (
	"fmt"
	"strings"
)

func main() {

	// NUMBERS
	var x int
	fmt.Println(x)

	var y = 4 // Note the type is inferred
	fmt.Println(y)

	z := 36 // Note the type is inferred
	fmt.Println(z)

	// STRINGS
	var name string
	name = "Mark"
	fmt.Println(name)

	lastName := "Smith"
	//lastName := "Jones" // This is an error: the variable is already declared
	lastName = "Jones" // This is fine
	fmt.Println(lastName, "Length:", len(lastName))

	fmt.Println(strings.Replace(lastName, "J", "B", 1))
	fmt.Println(strings.Split(lastName, ""))

	// ARRAYS
	var names [3]string
	names[0] = "Mark"
	names[1] = "John"
	names[2] = "Jane"
	fmt.Println(names)

	lastNames := [3]string{"Smith", "Jones", "Doe"}
	fmt.Println(lastNames)

	phrase := []string{"My", "name"}
	phrase = append(phrase, "is", "Mark")
	fmt.Println(phrase, "Length:", len(phrase))

	// STRUCTS
	p1 := Person{"Mark", 36}
	fmt.Println(p1)
	fmt.Println(p1.name, p1.age)
	p1.age = 25
	fmt.Println(p1)
	p1.Info()

}

func swap(m1, m2 *int) {
	var temp int
	temp = *m1 // temp holds the value of m1
	*m1 = *m2  // m1 now holds the value of m2
	*m2 = temp // m2 now holds the value of temp (which held the value of m1)
}

type Person struct {
	name string
	age  int
}

func (p Person) Info() {
	fmt.Println(fmt.Sprintf("%v has %v years", p.name, p.age))
}
