package main

import "fmt"

func main() {
	var no int = 10
	var intPtr *int = &no
	fmt.Println(intPtr)

	//deferencing
	fmt.Println(*intPtr)
	//no == *(&no)

	fmt.Println("Before incrementing, no = ", no)
	//increment(no)
	increment(&no)
	fmt.Println("After incrementing, no = ", no)

	var n1, n2 = 10, 20

	fmt.Println("Before swap, n1 = ", n1, " n2 = ", n2)
	swap( /* .... */ )
	fmt.Println("After swap, n1 = ", n1, " n2 = ", n2)
}

/*
func increment(x int) {
	x += 1
}
*/

func increment(x *int) {
	*x += 1
}

func swap( /* ... */ ) {
	/* ... */
}
