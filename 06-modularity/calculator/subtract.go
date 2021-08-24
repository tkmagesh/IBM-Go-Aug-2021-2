package calculator

func Subtract(x, y int) int {
	operationCount++
	return x - y
}

func GetOperationCount() int {
	return operationCount
}
