package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Variable
	var age int = 22

	var name = "Dung"

	var count int

	point := 5.5

	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(count) //default value = 0
	fmt.Println(point)

	point = 7.5

	fmt.Println(point)

	var x, y = 1, 4.5

	fmt.Println(x)
	fmt.Println(y)

	i, j := 1, 2

	fmt.Println(i)
	fmt.Println(j)

	// const PI float32 = 3.14
	const PI = 3.14

	fmt.Println(PI)

	var isValid bool
	fmt.Println(isValid) // default value = false

	var fullName string

	fmt.Println(fullName)       // default value = ""
	fmt.Println(fullName == "") // true

	var id, userId int16

	fmt.Println(id, userId)

	var c rune = 'a'

	fmt.Println(c)
	fmt.Printf("%c", c)
	fmt.Printf("%T", c)
	fmt.Printf("%p", &isValid)
	fmt.Println()
	fmt.Printf("%t", isValid)
	println()
	println("------------------------")

	str := "F"
	b, _ := strconv.ParseBool(str)
	fmt.Println(b)
	fmt.Printf("bool from bool: %v\n", b)
	fmt.Println()
}
