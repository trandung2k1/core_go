package main

import "fmt"

func main() {
	//==, !=, > , <, >=, <=, &&, ||, !
	var num1 int = 1
	var num2 int = 2

	result := num1 == num2
	fmt.Println(result)

	// if else
	number := 1
	if number > 0 {
		fmt.Printf("%d is a positive number\n", number)
	}
	fmt.Println("Out of the loop")

	if number > 0 {
		fmt.Printf("%d is a positive number\n", number)
	} else {
		fmt.Printf("%d is a negative number\n", number)
	}
	if num1 >= num2 {
		if num1 == num2 {
			fmt.Printf("Result: %d == %d", num1, num2)
		} else {
			fmt.Printf("Result: %d > %d", num1, num2)
		}
	} else {
		fmt.Printf("Result: %d < %d", num1, num2)
	}
	fmt.Println(number)
	num := 1
	switch num {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)

	default:
		fmt.Println("Invalid")
	}
	//Rs: 1

	switch num {
	case 1:
		fmt.Println(1)
		fallthrough
	case 2:
		fmt.Println(2)

	default:
		fmt.Println("Invalid")
	}
	//Rs: 1 , 2

	dayOfWeek := "Sunday"
	switch dayOfWeek {
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Weekday")

	default:
		fmt.Println("Invalid day")
	}

	numberOfDays := 28
	switch {

	case 28 == numberOfDays:
		fmt.Println("It's February")

	default:
		fmt.Println("Not February")
	}

}
