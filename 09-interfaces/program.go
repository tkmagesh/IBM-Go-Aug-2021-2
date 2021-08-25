package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func (c Circle) Perimeter() float64 {
	return c.Radius * 2 * math.Pi
}

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Height + 2*r.Width
}

//utility functions

type ShapeWithArea interface {
	Area() float64
}

func PrintArea(shaprWithArea ShapeWithArea) {
	fmt.Println("Area = ", shaprWithArea.Area())
}

type ShapeWithPerimeter interface {
	Perimeter() float64
}

func PrintPerimeter(shapeWithPerimeter ShapeWithPerimeter) {
	fmt.Println("Perimeter = ", shapeWithPerimeter.Perimeter())
}

//interface composition
type Dimension interface {
	ShapeWithArea
	ShapeWithPerimeter
}

func PrintDimensions(d Dimension) {
	fmt.Println("Dimensions")
	fmt.Println("Area = ", d.Area())
	fmt.Println("Permeter = ", d.Perimeter())
}

func main() {
	c := Circle{Radius: 10}
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintDimensions(c)

	r := Rectangle{Height: 10, Width: 12}
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintDimensions(r)
}
