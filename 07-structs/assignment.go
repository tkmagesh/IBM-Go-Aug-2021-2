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
		{105, "Pen", 50, 5, "Stationary"},
		{107, "Pencil", 100, 2, "Stationary"},
		{103, "Marker", 50, 20, "Utencil"},
		{102, "Stove", 5, 5000, "Utencil"},
		{101, "Kettle", 10, 2500, "Utencil"},
		{104, "Scribble Pad", 20, 20, "Stationary"},
	}
	fmt.Println("Initial List")
	fmt.Println(FormatProducts(products))

	fmt.Println("IndexOf")
	marker := Product{103, "Marker", 50, 20, "Utencil"}
	fmt.Println("IndexOf marker => ", IndexOf(products, marker))
	fmt.Println("")

	fmt.Println("Includes")
	kettle := Product{101, "Kettle", 10, 2500, "Utencil"}
	fmt.Println("Includes Kettle ? => ", Includes(products, kettle))
	fmt.Println("")

	fmt.Println("Filter : Costly Products")
	costlyProductPredicate := func(product Product) bool {
		return product.Cost > 100
	}
	costlyProducts := Filter(products, costlyProductPredicate)
	fmt.Println(FormatProducts(costlyProducts))

	fmt.Println("Filter : Stationary Products")
	stationaryProductPredicate := func(product Product) bool {
		return product.Category == "Stationary"
	}
	stationaryProducts := Filter(products, stationaryProductPredicate)
	fmt.Println(FormatProducts(stationaryProducts))

	fmt.Println("Any")
	anyCostlyProducts := Any(products, costlyProductPredicate)
	fmt.Println("Are there any costly products? : ", anyCostlyProducts)
	fmt.Println()

	fmt.Println("All")
	allCostlyProducts := All(products, costlyProductPredicate)
	fmt.Println("Are all products  costly products? : ", allCostlyProducts)

}

func FormatProduct(product Product) string {
	return fmt.Sprintf("Id = %d, Name = %s, Units = %d, Cost = %v, Category = %s", product.Id, product.Name, product.Units, product.Cost, product.Category)
}

func FormatProducts(products []Product) string {
	var s string
	for _, p := range products {
		s += fmt.Sprintf("%v\n", FormatProduct(p))
	}
	return s
}

//IndexOf => returns the index of the given product in the products collections (-1 if it doesnt exist)
func IndexOf(products []Product, product Product) int {
	for i, p := range products {
		if p == product {
			return i
		}
	}
	return -1
}

//Includes => return true/false based on the existence of the given product in the products collection
func Includes(products []Product, product Product) bool {
	return IndexOf(products, product) != -1
}

//Filter => returns a new collection of products that match the given criteria
//use cases
//filter costly products (cost > 100)
/*
func FilterCostlyProducts(products []Product) []Product {
	var filtered []Product
	for _, p := range products {
		if p.Cost > 100 {
			filtered = append(filtered, p)
		}
	}
	return filtered
}
*/

//filter products by category (category == "Stationary")
/*
func FilterByCategory(products []Product, category string) []Product {
	var filtered []Product
	for _, p := range products {
		if p.Category == category {
			filtered = append(filtered, p)
		}
	}
	return filtered
}
*/

func Filter(products []Product, predicate func(Product) bool) []Product {
	var filtered []Product
	for _, p := range products {
		if predicate(p) {
			filtered = append(filtered, p)
		}
	}
	return filtered
}

//filter products by category and cost (category == "Stationary" && cost > 100)

//Any => return true/false based on the existence of a product in the collection that satisfies the given criteria
//use cases
//Are there any costly products (cost > 100)
//Are there any "Stationary" products

func Any(products []Product, predicate func(Product) bool) bool {
	for _, p := range products {
		if predicate(p) {
			return true
		}
	}
	return false
}

//All => return true/false if all the products in the collection satisfies the given criteria
//use cases
//Are all the products costly products (cost > 100)
//Are all the products "Stationary" products

func All(products []Product, predicate func(Product) bool) bool {
	for _, p := range products {
		if !predicate(p) {
			return false
		}
	}
	return true
}
