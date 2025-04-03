package main

import "fmt"

func main() {

	f := true
	flag := &f

	if flag == nil {
		fmt.Println("Flag is nil")
	} else if *flag {
		fmt.Println("Flag is true")
	} else {
		fmt.Println("Flag is false")
	}

	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			fmt.Println(i, " is even")
		} else if i == 3 {
			continue
		} else {
			fmt.Println(i, " is odd")
		}
	}

	arr := []string{"John", "Mike", "Jane"}
	for i, name := range arr {
		fmt.Println(i, name)
	}

	mymap := map[string]string{"name": "John", "age": "30"}
	for key, value := range mymap {
		fmt.Println(key, value)
	}

	day := "Friday"
	switch day {
	case "Monday":
		fmt.Println("It's Monday")
	case "Tuesday":
		fmt.Println("It's Tuesday")
	case "Friday":
		fmt.Println("It's Friday, let's have a party")
	default:
		fmt.Println("It's another day")
	}

	year := 2025
	switch {
	case year%4 == 0 && year%100 != 0:
		fmt.Println("It's a leap year")
	case year%400 == 0:
		fmt.Println("It's a leap year")
	default:
		fmt.Println("It's not a leap year")
	}
}
