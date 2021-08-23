package main

import "fmt"

//unused variables are allowed at the package level (but not at the function)
var x = 100

func main() {
	//programming constucts -> if, for, switch case

	/*
		no := 21
		if no%2 == 0 {
			fmt.Println(no, " is an even number")
		} else {
			fmt.Println(no, " is an odd number")
		}
	*/

	var x = 100

	_ = x

	if no := 21; no%2 == 0 {
		fmt.Println(no, " is an even number")
	} else {
		fmt.Println(no, " is an odd number")
	}

}
