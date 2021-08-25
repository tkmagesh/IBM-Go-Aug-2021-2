package main

import (
	"fmt"
)

func main() {
	go print("Hello")
	print("World")

	//time.Sleep(1 * time.Second)

	var input string
	fmt.Scanln(&input)
}

func print(str string) {
	fmt.Println(str)
}
