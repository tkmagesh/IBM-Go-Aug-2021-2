package main

import "fmt"

func main() {
	hello := getFn()
	hello()

	add := wrap(add)
	subtract := wrap(subtract)
	add(100, 200)
	subtract(100, 200)
}

func getFn() func() {
	return func() {
		println("Hello")
	}
}

func wrap(oper func(int, int) int) func(int, int) {
	return func(a, b int) {
		fmt.Println("Before invocation")
		fmt.Println(oper(a, b))
		fmt.Println("After invocation")
	}
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}
