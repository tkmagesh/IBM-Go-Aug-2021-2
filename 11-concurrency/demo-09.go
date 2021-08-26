package main

import "time"

func main() {
	print("Hello", time.Second*2)
	print("World", time.Second*1)
}

func print(str string, delay time.Duration /* other parameter */) {
	for i := 0; i < 5; i++ {
		time.Sleep(delay)
		println(str)
	}
}

/*
Desired Output
Hello
World
Hello
World
Hello
World
Hello
World
Hello
World
*/
