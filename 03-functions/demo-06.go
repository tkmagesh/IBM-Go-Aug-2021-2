package main

import "fmt"

func main() {
	inc := getIncrement()
	fmt.Println(inc()) // => 1
	fmt.Println(inc()) // => 2
	//reset()

	fmt.Println(inc()) // => 3
	fmt.Println(inc()) // => 4
}

func getIncrement() func() int { //step 1
	var counter = 0    //step 2
	fn := func() int { //step 3
		counter += 1 //step 4
		return counter
	}
	return fn //step 5
}
