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

	//var x = 100

	if no := 21; no%2 == 0 {
		fmt.Println(no, " is an even number")
	} else {
		fmt.Println(no, " is an odd number")
	}

	//for construct
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//using 'for' like a while construct
	numSum := 1
	for numSum < 100 {
		numSum += numSum
	}
	fmt.Println(numSum)

	//infinite loop
	x := 1
	for {
		x += x
		if x > 100 {
			break
		}
	}
	fmt.Println(x)
}
