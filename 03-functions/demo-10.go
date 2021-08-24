package main

import "fmt"

func main() {

	defer fmt.Println("Deferred from main")
	fmt.Println("main invoked")
	f1()
}

func f1() {

	defer fmt.Println("Deferred(1) from f1")
	defer fmt.Println("Deferred(2) from f1")
	defer fmt.Println("Deferred(3) from f1")
	fmt.Println("f1 invoked")
	f2()
}

func f2() {

	defer fmt.Println("Deferred from f2")
	fmt.Println("f2 invoked")
}

func processFile() {
	//open the file
	defer func() {
		//close the file
	}()

	//read the file
	// if err != nil {
	//		return
	// }

	//while not eof

	//parse the line
	// if err != nil {
	//		return
	// }

	//process the data
	// if err != nil {
	//		return
	// }

	//read the file
	// if err != nil {
	//		return
	// }

	//close the file

	//return the result
}
