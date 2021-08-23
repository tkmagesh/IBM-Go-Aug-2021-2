package main

import "fmt"

func main() {
	hello := getFn()
	hello()

	add := wrap("add", add)
	subtract := wrap("subtract", subtract)
	add(100, 200)
	subtract(100, 200)
}

func getFn() func() {
	return func() {
		println("Hello")
	}
}

func wrap(fnName string, oper func(int, int) int) func(int, int) {
	return func(a, b int) {
		fmt.Printf("Before invoking %s function\n", fnName)
		fmt.Println(oper(a, b))
		fmt.Printf("After invoking %s function\n", fnName)
	}
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}
