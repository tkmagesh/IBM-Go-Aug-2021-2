package main

import "fmt"

func main() {

	/* wrappedAdd(10, 20)
	wrappedSubtract(10, 20) */
	wrappedInvoke(10, 20, add)
	wrappedInvoke(10, 20, subtract)
}

func wrappedAdd(a int, b int) {
	fmt.Println("Before invocation")
	fmt.Println(add(a, b))
	fmt.Println("After invocation")
}

func wrappedSubtract(a, b int) {
	fmt.Println("Before invocation")
	fmt.Println(subtract(a, b))
	fmt.Println("After invocation")
}

func wrappedInvoke(a, b int, oper func(int, int) int) {
	fmt.Println("Before invocation")
	fmt.Println(oper(a, b))
	fmt.Println("After invocation")
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}
