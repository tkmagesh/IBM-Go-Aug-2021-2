package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Units    int
	Cost     float64
	Category string
}

func main() {

	var pen Product = Product{Id: 100, Name: "Pen", Units: 10, Cost: 5.00, Category: "Stationary"}
	fmt.Println(pen)
	fmt.Printf("Product %s costs %v\n", pen.Name, pen.Cost)
	//applyDiscount(&pen, 10)
	pen.applyDiscount(10)
	fmt.Println("After applying 10% discount")
	fmt.Println(pen)
}

func (p *Product) applyDiscount(discount int) {
	p.Cost = p.Cost - (p.Cost * float64(discount) / 100)
}
