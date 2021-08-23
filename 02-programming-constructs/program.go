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

	//switch case construct
	/*
		no := 20
		remainder := no % 2
		switch remainder {
		case 0:
			fmt.Println(no, " is an even number")
		case 1:
			fmt.Println(no, " is an odd number")
		}
	*/

	/*
		switch no := 20; {
		case no%2 == 0:
			fmt.Println(no, " is an even number")
		case no%2 == 1:
			fmt.Println(no, " is an odd number")
		}
	*/

	/*
		rank by score
		score 0 - 3 -> "Terrible"
		score 4 - 7 -> "Not Bad"
		score 8 - 9 -> "Good"
		score 10 -> "Excellent"
		otherwise -> "Unknown score"
	*/

	score := 6
	switch score {
	case 0, 1, 2, 3:
		fmt.Println("Terrible")
	case 4, 5, 6, 7:
		fmt.Println("Not Bad")
	case 8, 9:
		fmt.Println("Good")
	case 10:
		fmt.Println("Excellent")
	default:
		fmt.Println("Unknown score")
	}
}
