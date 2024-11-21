package main

import (
	"fmt"
)

type Info struct {
	name string
	surname string
	age int8
	money float32
}

func main() {
	var info1 Info
	var info2 Info

	info1.name = "Aydın"
	info1.surname = "Demirci"
	info1.age = 124
	info1.money = 12.3

	info2.name = "John"
	info2.surname = "Doe"
	info2.age = 45
	info2.money = 134223.332

	fmt.Println("Name:", info1.name)
	fmt.Println("Surname:", info1.surname)
	fmt.Println("Age:", info1.age)
	fmt.Println("Money:", info1.money)


	fmt.Println("Name:", info2.name)
	fmt.Println("Surname:", info2.surname)
	fmt.Println("Age:", info2.age)
	fmt.Println("Money:", info2.money)
}