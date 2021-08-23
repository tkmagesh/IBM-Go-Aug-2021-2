package main

import "fmt"

func main() {
	n1, n2 := 10, 20
	result := n1 + n2
	//fmt.Println("Sum of ", n1, " and ", n2, " is ", result)
	fmt.Printf("Sum of %d and %d is %d\n", n1, n2, result)

	fmt.Println("Enter you name :")
	var input string
	fmt.Scanln(&input)
	fmt.Printf("Hello %s, Welcome to golang!\n", input)

	fmt.Println("Enter any number :")
	var num int
	fmt.Scanln(&num)
	fmt.Println("You entered ", num)
}
