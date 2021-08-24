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

func NewPerishableProduct(id int, name string, units int, cost float64, category string, expiry string) *PerishableProduct {
	return &PerishableProduct{
		Product: Product{id, name, units, cost, category},
		Expiry:  expiry,
	}
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

	/*
		grapes := PerishableProduct{
			Product: Product{101, "Grapes", 100, 1.99, "Fruit"},
			Expiry:  "3 Days",
		}
	*/

	grapes := NewPerishableProduct(101, "Grapes", 100, 1.99, "Fruit", "3 Days")

	fmt.Println(grapes)
	//fmt.Printf("Product %s costs %v\n", grapes.Product.Name, grapes.Product.Cost)
	fmt.Printf("Product %s costs %v\n", grapes.Name, grapes.Cost)

	//TO BE FIXED
	applyDiscount(&grapes.Product, 10)

	fmt.Printf("Product %s costs %v\n", grapes.Name, grapes.Cost)
}

func applyDiscount(p *Product, discount int) {
	p.Cost = p.Cost - (p.Cost * float64(discount) / 100)
}
