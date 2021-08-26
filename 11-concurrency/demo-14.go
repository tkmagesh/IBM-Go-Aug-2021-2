package main

import (
	"fmt"
	"time"
)

/*
func main() {
	fibCh := fibonacci()
	for no := range fibCh {
		fmt.Println(no)
	}
	fmt.Println("Done")
}

func fibonacci() chan int {
	fibCh := make(chan int)
	go func() {
		x, y := 0, 1
		for i := 0; i < 10; i++ {
			fibCh <- x
			x, y = y, x+y
		}
		close(fibCh)
	}()
	return fibCh
}
*/

func main() {
	fibCh := make(chan int)
	doneCh := make(chan bool)
	go fibonacci(fibCh, doneCh)
	go func() {
		for no := range fibCh {
			fmt.Println(no)
		}
	}()
	fmt.Println("Hit Enter to stop!")
	var input string
	fmt.Scanln(&input)
	doneCh <- true
	fmt.Println("[main] Done")
}

func fibonacci(fibCh chan int, doneCh chan bool) {
	x, y := 0, 1
	for {
		select {
		case fibCh <- x:
			time.Sleep(500 * time.Millisecond)
			x, y = y, x+y
		case <-doneCh:
			fmt.Println("Finished")
			close(fibCh)
			return
		}
	}
}
