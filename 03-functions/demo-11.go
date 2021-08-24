package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from a panic in main:", r)
		}
	}()
	fmt.Println(divide(100, 7))
	fmt.Println(divide(100, 0))
}

func divide(x, y int) int {
	if y == 0 {
		panic("divide by zero")
	}
	return x / y
}
