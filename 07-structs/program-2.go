package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Units    int
	Cost     float64
	Category string
}

/*
type PerishableProduct struct {
	Prod   Product
	Expiry string
}
*/

type PerishableProduct struct {
	Product
	Expiry string
}

func main() {
	/*
		grapes := PerishableProduct{
			Prod:   Product{101, "Grapes", 100, 1.99, "Fruit"},
			Expiry: "3 Days",
		}

		fmt.Println(grapes)
		fmt.Printf("Product %s costs %v\n", grapes.Prod.Name, grapes.Prod.Cost)
	*/

	/*
		grapes := PerishableProduct{
			Product{101, "Grapes", 100, 1.99, "Fruit"},
			"3 Days",
		}
	*/

	grapes := PerishableProduct{
		Product: Product{101, "Grapes", 100, 1.99, "Fruit"},
		Expiry:  "3 Days",
	}

	fmt.Println(grapes)
	//fmt.Printf("Product %s costs %v\n", grapes.Product.Name, grapes.Product.Cost)
	fmt.Printf("Product %s costs %v\n", grapes.Name, grapes.Cost)
}
