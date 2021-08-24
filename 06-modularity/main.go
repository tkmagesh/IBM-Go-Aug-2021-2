package main

import (
	"fmt"
	"modularity_app/calculator"
	"modularity_app/calculator/utils"
	sp "modularity_app/spinner"
)

func main() {
	fmt.Println(calculator.Add(10, 20))
	fmt.Println(calculator.Subtract(100, 200))
	fmt.Println(calculator.GetOperationCount())

	fmt.Println(utils.IsPrime(71))

	fmt.Println("Spinner")
	fmt.Println(sp.Increment())
	fmt.Println(sp.Increment())
	fmt.Println(sp.Increment())
	fmt.Println(sp.Increment())

	fmt.Println(sp.Decrement())
	fmt.Println(sp.Decrement())
	fmt.Println(sp.Decrement())
	fmt.Println(sp.Decrement())
	fmt.Println(sp.Decrement())
}
