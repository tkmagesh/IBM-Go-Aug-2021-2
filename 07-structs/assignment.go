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
	products := []Product{
		{105, "Pen", 5, 50, "Stationary"},
		{107, "Pencil", 2, 100, "Stationary"},
		{103, "Marker", 50, 20, "Utencil"},
		{102, "Stove", 5000, 5, "Utencil"},
		{101, "Kettle", 2500, 10, "Utencil"},
		{104, "Scribble Pad", 20, 20, "Stationary"},
	}

	for _, p := range products {
		fmt.Println(p)
	}
}

//IndexOf => returns the index of the given product in the products collections (-1 if it doesnt exist)

//Includes => return true/false based on the existence of the given product in the products collection

//Filter => returns a new collection of products that match the given criteria
//use cases
//filter costly products (cost > 100)
//filter products by category (category == "Stationary")
//filter products by category and cost (category == "Stationary" && cost > 100)

//Any => return true/false based on the existence of a product in the collection that satisfies the given criteria
//use cases
//Are there any costly products (cost > 100)
//Are there any "Stationary" products

//All => return true/false if all the products in the collection satisfies the given criteria
//use cases
//Are all the products costly products (cost > 100)
//Are all the products "Stationary" products
