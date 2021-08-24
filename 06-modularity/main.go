package main

import (
	"fmt"
	"modularity_app/calculator"
	"modularity_app/calculator/utils"
	sp "modularity_app/spinner"

	"github.com/fatih/color"
)

func main() {
	fmt.Println(calculator.Add(10, 20))
	fmt.Println(calculator.Subtract(100, 200))
	fmt.Println(calculator.GetOperationCount())

	fmt.Println(utils.IsPrime(71))

	color.Red("Spinner")
	color.Yellow("%d\n", sp.Increment())
	color.Yellow("%d\n", sp.Increment())
	color.Yellow("%d\n", sp.Increment())
	color.Yellow("%d\n", sp.Increment())

	color.Green("%d\n", sp.Decrement())
	color.Green("%d\n", sp.Decrement())
	color.Green("%d\n", sp.Decrement())
	color.Green("%d\n", sp.Decrement())
	color.Green("%d\n", sp.Decrement())

}
