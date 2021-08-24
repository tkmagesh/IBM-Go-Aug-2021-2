package main

import (
	"fmt"
	"os"
)

func main() {
	for {
		userChoice := getUserChoice()
		switch userChoice {
		case 1, 2, 3, 4:
			processUserChoice(userChoice)
		case 5:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func getOperands() (int, int) {
	var n1, n2 int
	fmt.Println("Enter two numbers:")
	fmt.Scanf("%d %d", &n1, &n2)
	return n1, n2
}

func processUserChoice(userChoice int) {
	var result int
	n1, n2 := getOperands()
	switch userChoice {
	case 1:
		result = add(n1, n2)
	case 2:
		result = subtract(n1, n2)
	case 3:
		result = multiply(n1, n2)
	case 4:
		result = divide(n1, n2)
	}
	fmt.Println("Result = ", result)
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}

func getUserChoice() int {
	var userChoice int
	fmt.Println("Menu")
	fmt.Println("================================================================")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Exit")
	fmt.Println("Enter your choice:")
	fmt.Scanf("%d", &userChoice)
	return userChoice
}
