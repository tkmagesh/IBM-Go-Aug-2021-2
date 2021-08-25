package main

import (
	"fmt"
	"sync"
)

var wg *sync.WaitGroup = &sync.WaitGroup{}

func main() {
	wg.Add(2)

	go print("Hello")
	go print("World")

	//time.Sleep(1 * time.Second)

	/*
		var input string
		fmt.Scanln(&input)
	*/
	wg.Wait()
	fmt.Println("Exiting from main")
}

func print(str string) {
	fmt.Println(str)
	wg.Done()
}
