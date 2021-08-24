package calculator

var operationCount int

func Add(x, y int) int {
	operationCount++
	return x + y
}
