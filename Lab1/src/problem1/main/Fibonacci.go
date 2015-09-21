package main

import (
	"fmt"
)

/*
Function which returns N'th number in fibonacci series
*/
func Fibonacci(number int) int64 {

	if number == 0 {
		return 0
	}

	if number == 1 {
		return 1
	}

	return Fibonacci(number-1) + Fibonacci(number-2)
}

func main() {
	var number int
	fmt.Println("Enter a positive number:")
	fmt.Scanf("%d", &number)
	fmt.Println("Fibonacci of ", number, " is ", Fibonacci(number))
}
