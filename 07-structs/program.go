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
	/*
		var pen Product = Product{}
		pen.Id = 100
		pen.Name = "Pen"
		pen.Units = 10
		pen.Cost = 5.00
		pen.Category = "Stationary"
	*/
	//var pen Product = Product{100, "Pen", 10, 5.00, "Stationary"}
	var pen Product = Product{Id: 100, Name: "Pen", Units: 10, Cost: 5.00, Category: "Stationary"}
	fmt.Println(pen)
	fmt.Printf("Product %s costs %v\n", pen.Name, pen.Cost)
	applyDiscount(&pen, 10) //apply 10% discount on the product cost
	fmt.Println("After applying 10% discount")
	fmt.Println(pen)
}

func applyDiscount(p *Product, discount int) {
	p.Cost = p.Cost - (p.Cost * float64(discount) / 100)
}
