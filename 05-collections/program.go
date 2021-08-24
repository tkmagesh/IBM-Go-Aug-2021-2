package main

import (
	"fmt"
	"strings"
)

func main() {
	//Arrays
	//var nos [5]int
	//var nos [5]int = [5]int{3, 1, 4, 2, 5}
	var nos [5]int = [...]int{3, 1, 4, 2, 5}
	//fmt.Println(nos)

	//iterating over arrays
	/*
		for i := 0; i < len(nos); i++ {
			fmt.Println(nos[i])
		}
	*/

	//using the range construct
	/*
		for idx, no := range nos {
			fmt.Println(idx, no)
		}
	*/

	for _, no := range nos {
		fmt.Println(no)
	}

	//copying an array
	newNos := nos
	nos[0] = 200
	fmt.Println(nos, newNos)

	//Slice
	//var names []string
	var names []string = []string{"Magesh", "Ramesh", "Ganesh"}
	fmt.Println(names)

	//names = append(names, "Rajesh")
	//names = append(names, "Rajesh", "Suresh")

	newNames := []string{"Rajesh", "Suresh"}
	//names = append(names, newNames[0], newNames[1])
	names = append(names, newNames...)
	fmt.Println(names)

	//slicing
	/*
		[lo : hi] => from lo to hi-1
		[lo : ] => from lo to the end
		[ : hi] => from the 0 to hi-1
		[ lo : lo ] => []
		[lo : lo+1] => [lo]
		[:] => copy of the slice
	*/

	fmt.Println("names[1:3] => ", names[1:3])
	fmt.Println("names[:3] => ", names[:3])
	fmt.Println("names[3:] => ", names[3:])

	//maps
	var cityRanks map[string]int = map[string]int{
		"Bengaluru": 4,
		"Mysuru":    2,
		"Udupi":     1,
		"Mangaluru": 3,
	}
	fmt.Println(cityRanks, len(cityRanks))

	var rankOfMysuru = cityRanks["Mysuru"]
	fmt.Println("Rank of Mysuru => ", rankOfMysuru)

	//adding a new key/value pair
	cityRanks["Pune"] = 5
	fmt.Println(cityRanks)

	//check if a key exists or not
	if _, ok := cityRanks["Chennai"]; ok {
		fmt.Println("Chennai exists in the map")
	} else {
		fmt.Println("Chennai does not exist in the map")
	}

	//iterating through a map
	for city, rank := range cityRanks {
		fmt.Println(city, rank)
	}

	s := "Officia enim irure aute eiusmod mollit cupidatat veniam duis"
	fmt.Println(strings.Split(s, " "))
}
