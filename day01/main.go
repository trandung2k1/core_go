package main

import (
	"fmt"
	"time"
)

func main() {
	sayHi()
	total := sum(1, 2)
	fmt.Println(total)

	// Goroutine
	go sayHello("Viet")
	// normal function
	sayHello("Nam")
	time.Sleep(time.Second)
}

func sayHi() {
	fmt.Println("Say Hi")
}

func sum(a int, b int) int {
	return a + b
}

func variable() {
	var msg string = "Hello"
	msg = "Xin chao"
	fmt.Println(msg)
	fmt.Printf("%T", msg)
	fmt.Println()
	fmt.Printf("%V", msg)
	fullName := "Tran Van Dung"
	println(fullName)
	var age = 10
	var point float32 = 5.5

	fmt.Println(age)
	println(point)
	fmt.Printf("%T", age)
	fmt.Println()
	fmt.Printf("%T", point)
}

func sayHello(name string) {
	for i := 0; i <= 5; i++ {
		fmt.Printf("Hello %s\n", name)
	}
}
