package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(sum(10))                                     //=> 10
	fmt.Println(sum(10, 20))                                 //=> 30
	fmt.Println(sum(10, 20, 30, 40, 50))                     //=> 150
	fmt.Println(sum(10, 20, "30", 40, 50))                   //=> 150
	fmt.Println(sum(10, 20, "abc", 40, 50))                  //=> 120
	fmt.Println(sum(10, 20, "abc", []int{40, 50}))           //=> 120
	fmt.Println(sum(10, 20, "abc", []interface{}{40, "50"})) //=> 120
}

func sum(nos ...interface{}) int {
	result := 0
	for _, no := range nos {
		switch v := no.(type) {
		case int:
			result += v
		case string:
			if val, err := strconv.Atoi(v); err == nil {
				result += val
			}
		case []int:
			intfList := make([]interface{}, len(v))
			for idx, x := range v {
				intfList[idx] = x
			}
			result += sum(intfList...)
		case []interface{}:
			result += sum(v...)
		}

	}
	return result
}
