package main

import "fmt"

func main() {
	data := []int{4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55}

	dataSet1 := data[:len(data)/2]
	dataSet2 := data[len(data)/2:]
	resultCh := make(chan int)
	go func() {
		resultCh <- sum(dataSet1...)
	}()
	go func() {
		resultCh <- sum(dataSet2...)
	}()
	finalResult := <-resultCh + <-resultCh
	fmt.Println(finalResult)
}

func sum(nos ...int) int {
	result := 0
	for _, no := range nos {
		result += no
	}
	return result
}
