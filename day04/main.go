package main

import "fmt"

func main() {

	var num1 int64 = 10
	var num2 int64 = 20

	fmt.Println(num1 + num2)
	fmt.Println(num1 - num2)
	fmt.Println(num1 * num2)
	fmt.Println(num1 / num2)
	fmt.Println(num1 % num2)

	// ++, --, +=, -=, *=, /=, %=, ==, !=, >, <, >=, <=, &&, ||, !

	var floatValue float32 = 9.8
	var intValue int = int(floatValue)
	fmt.Println(intValue)

	var age int = 22
	var floatAge float32 = float32(age)
	fmt.Printf("%T", floatAge)
	fmt.Printf("\n%f", floatAge)

	var point = 5.5
	fmt.Printf("\n%g", point)
}
