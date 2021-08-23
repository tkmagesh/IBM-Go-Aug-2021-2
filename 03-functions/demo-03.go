package main

import "fmt"

func main() {
	fn := func() {
		fmt.Println("fn is invoked")
	}
	fn()

	/*
		add := func(x, y int) int {
			return x + y
		}
	*/
	var add func(int, int) int = func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 2))

	//anonymous functions
	func() {
		fmt.Println("anonymous function")
	}()

	quotient, remainder := func(x, y int) (int, int) {
		return x / y, x % y
	}(100, 7)
	fmt.Println(quotient, remainder)

}
