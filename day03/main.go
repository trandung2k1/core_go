package main

import (
	"fmt"
)

func main() {

	//Print
	name := "Dung"
	fmt.Print("Name : ", name)
	fmt.Print("\nName : ", name)

	//Println

	println()
	currentSalary := 50000
	fmt.Println(currentSalary)

	//Printf

	currentAge := 22
	fmt.Printf("Age = %d", currentAge)
	println()

	//int: %d
	//string: %s
	//bool: %t
	//float: %g

	msg := "Hello"
	println(msg)

	print(currentAge)
	println()

	//Take input

	//Scan
	var fullName string
	var age int
	// fmt.Printf("Enter your name: ")
	// fmt.Scan(&fullName)

	// fmt.Print("Enter your name and age: ")
	// fmt.Scan(&fullName, &age)
	// fmt.Printf("Fullname: %s\nAge: %d", fullName, age)

	//Scanln
	// fmt.Println("Enter your name and age: ")
	// fmt.Scanln(&fullName, &age)
	// fmt.Printf("Fullname: %s\nAge: %d", fullName, age)

	//Scanf
	fmt.Println("Enter your name ")
	fmt.Scanf("%s %d", &fullName, &age)
	fmt.Printf("Fullname: %s\nAge: %d", fullName, age)

	//Comment
	//One line

	/*
		Multiple line
	*/
}
