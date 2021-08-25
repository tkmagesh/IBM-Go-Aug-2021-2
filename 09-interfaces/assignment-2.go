package main

import "fmt"

func main() {
	fmt.Println(sum(10)) //=> 10
	fmt.Println(sum(10, 20)) //=> 30
	fmt.Println(sum(10, 20, 30, 40, 50)) //=> 150
	fmt.Println(sum(10, 20, "30", 40, 50)) //=> 150
	fmt.Println(sum(10, 20, "abc", 40, 50)) //=> 120
	fmt.Println(sum(10, 20, "abc", []int{40, 50})) //=> 120
	fmt.Println(sum(10, 20, "abc", []interface{40, "50"})) //=> 120
}

func sum(nos ...int /* nos = slice of int */) int {
	result := 0
	for i := 0; i < len(nos); i++ {
		result += nos[i]
	}
	return result
}
