package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError error = errors.New("Invalid arguments. divide by zero")

func main() {
	fmt.Println(divide(100, 7))
	result, err := divide(100, 0)
	if err == DivideByZeroError {
		fmt.Println("Very bad attempt. Do not divide by zero")
		return
	}
	if err != nil {
		fmt.Println("something went wrong!")
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, DivideByZeroError
	}

	return x / y, nil
}
