package main

import "fmt"

func main() {
	fmt.Println(add(100, 200))
	//fmt.Println(divide(100, 7))
	q, r := divide(100, 7)
	fmt.Println(q, r)
}

func add(x int, y int) int {
	return x + y
}

/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

func divide(x, y int) (quotient int, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
