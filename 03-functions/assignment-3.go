package main

import "fmt"

func main(){
	isPrime := func() func(int) bool {
		prevNo := -1
		prevResult := false
		return func(no int) bool {
			if no == prevNo {
				return prevResult
			}
			fmt.Printf("Processing %d\n", no)
			prevNo = no
			prevResult = true
			for i := 2; i <= (no/2); i++ {
				if no % 2 == 0 {
					prevResult = false
					break
				}
			}
			return prevResult
		}
	}()

	fmt.Println(isPrime(10))
	fmt.Println(isPrime(10))
	fmt.Println(isPrime(10))
	fmt.Println(isPrime(11))
	fmt.Println(isPrime(11))
	fmt.Println(isPrime(10))
}

