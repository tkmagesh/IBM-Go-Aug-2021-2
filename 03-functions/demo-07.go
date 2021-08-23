package main

import "fmt"

func main() {
	inc, dec := spinner()
	fmt.Println(inc()) // => 1
	fmt.Println(inc()) // => 2
	//reset()

	fmt.Println(inc()) // => 3
	fmt.Println(inc()) // => 4
	fmt.Println(dec())
	fmt.Println(dec())
	fmt.Println(dec())
	fmt.Println(dec())
}

func spinner() (inc func() int, dec func() int) { //step 1
	var counter = 0    //step 2
	inc = func() int { //step 3
		counter += 1 //step 4
		return counter
	}
	dec = func() int {
		counter -= 1
		return counter
	}
	return
}
