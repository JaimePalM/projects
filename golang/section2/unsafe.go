package main

import "fmt"

func main_unsafe() {

	// POINTERS
	m1 := 15
	var ptr *int
	ptr = &m1
	fmt.Println(ptr)  // ptr holds the memory address of m1
	fmt.Println(*ptr) // *ptr holds the value of m1
	m2 := 20
	fmt.Println(m1, m2)
	swap(&m1, &m2)
	fmt.Println(m1, m2)

}
