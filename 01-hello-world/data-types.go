package main

import "fmt"

func main() {
	/*
		var msg string
		msg = "Hello World!"
	*/

	/*
		var msg string = "Hello World!"
	*/

	//type inference
	/*
		var msg = "Hello World!"
	*/

	//the following syntax is applicable only in function (not in package level variables)
	msg := "Hello World!"
	fmt.Println(msg)

	/*
		var label string
		var x int
		var y int
	*/

	/*
		var label string
		var x, y int
	*/

	/*
		var (
			label string
			x,y  int
		)
	*/

	/*
		var (
			label string = "Result = "
			x     int    = 100
			y     int    = 200
		)
	*/

	/*
		var (
			label = "Result = "
			x     = 100
			y     = 200
		)
	*/

	/*
		var (
			label = "Result = "
			x, y  = 100, 200
		)
	*/

	/*
		var (
			label, x, y = "Result = ", 100, 200
		)
	*/

	/*
		var label, x, y = "Result = ", 100, 200
	*/

	label, x, y := "Result = ", 100, 200
	fmt.Println(label, x+y)

}
