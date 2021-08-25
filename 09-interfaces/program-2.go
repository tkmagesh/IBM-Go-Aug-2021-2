package main

import "fmt"

type Product struct {
}

type Dummy struct {
}

func (d Dummy) M() {
	fmt.Println("Dummy method")
}

type I interface {
	M()
}

func main() {
	var x interface{}
	x = 100
	x = "This is a string"
	x = true
	x = struct{}{}
	x = []int{1, 2, 3}
	fmt.Println(x)

	var val interface{}
	val = 100
	if no, ok := val.(int); ok {
		y := no + 200
		fmt.Println(y)
	} else {
		fmt.Println("Not an int")
	}

	//val = "This is a string"
	//val = struct{}{}
	//val = Product{}
	val = Dummy{}
	switch v := val.(type) {
	case int:
		fmt.Printf("Double of val is %d\n", v*2)
	case string:
		fmt.Printf("Length of string %q is %d\n", v, len(v))
	case []int:
		fmt.Println("List of numbers => ", v)
	case struct{}:
		fmt.Println("Empty struct")
	case Product:
		fmt.Println("Instance of Product => ", v)
	case I:
		fmt.Println("implementation of I => ", v)
	default:
		fmt.Println("Unknown value")
	}

}
