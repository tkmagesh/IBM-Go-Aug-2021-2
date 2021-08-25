package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Units    int
	Cost     float64
	Category string
}

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

	grapes := NewPerishableProduct(101, "Grapes", 100, 1.99, "Fruit", "3 Days")

	fmt.Println(grapes)

	fmt.Printf("Product %s costs %v\n", grapes.Name, grapes.Cost)

	//TO BE FIXED
	//applyDiscount(&grapes.Product, 10)
	grapes.applyDiscount(10)

	fmt.Printf("Product %s costs %v\n", grapes.Name, grapes.Cost)
}

func (p *Product) applyDiscount(discount int) {
	p.Cost = p.Cost - (p.Cost * float64(discount) / 100)
}
