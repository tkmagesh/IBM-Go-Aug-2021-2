package main

import "fmt"

/* Product Type */
type Product struct {
	Id       int
	Name     string
	Units    int
	Cost     float64
	Category string
}

//Stringer interface implementation
func (product Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %s, Units = %d, Cost = %v, Category = %s", product.Id, product.Name, product.Units, product.Cost, product.Category)
}

/* Products Type */
type Products []Product

//Stringer interface implementation
func (products Products) String() string {
	var s string
	for _, p := range products {
		s += fmt.Sprintf("%v\n", p)
	}
	return s
}

func (products Products) IndexOf(product Product) int {
	for i, p := range products {
		if p == product {
			return i
		}
	}
	return -1
}

func (products Products) Includes(product Product) bool {
	return products.IndexOf(product) != -1
}

func (products Products) Filter(predicate func(Product) bool) Products {
	var filtered Products
	for _, p := range products {
		if predicate(p) {
			filtered = append(filtered, p)
		}
	}
	return filtered
}

func (products Products) Any(predicate func(Product) bool) bool {
	for _, p := range products {
		if predicate(p) {
			return true
		}
	}
	return false
}

func (products Products) All(predicate func(Product) bool) bool {
	for _, p := range products {
		if !predicate(p) {
			return false
		}
	}
	return true
}

func main() {
	products := Products{
		{105, "Pen", 50, 5, "Stationary"},
		{107, "Pencil", 100, 2, "Stationary"},
		{103, "Marker", 50, 20, "Utencil"},
		{102, "Stove", 5, 5000, "Utencil"},
		{101, "Kettle", 10, 2500, "Utencil"},
		{104, "Scribble Pad", 20, 20, "Stationary"},
	}
	fmt.Println("Initial List")
	fmt.Println(products)

	fmt.Println("IndexOf")
	marker := Product{103, "Marker", 50, 20, "Utencil"}
	fmt.Println("IndexOf marker => ", products.IndexOf(marker))
	fmt.Println("")

	fmt.Println("Includes")
	kettle := Product{101, "Kettle", 10, 2500, "Utencil"}
	fmt.Println("Includes Kettle ? => ", products.Includes(kettle))
	fmt.Println("")

	fmt.Println("Filter : Costly Products")
	costlyProductPredicate := func(product Product) bool {
		return product.Cost > 100
	}
	costlyProducts := products.Filter(costlyProductPredicate)
	fmt.Println(costlyProducts)

	fmt.Println("Filter : Stationary Products")
	stationaryProductPredicate := func(product Product) bool {
		return product.Category == "Stationary"
	}
	stationaryProducts := products.Filter(stationaryProductPredicate)
	fmt.Println(stationaryProducts)

	fmt.Println("Any")
	anyCostlyProducts := products.Any(costlyProductPredicate)
	fmt.Println("Are there any costly products? : ", anyCostlyProducts)
	fmt.Println()

	fmt.Println("All")
	allCostlyProducts := products.All(costlyProductPredicate)
	fmt.Println("Are all products  costly products? : ", allCostlyProducts)

}
