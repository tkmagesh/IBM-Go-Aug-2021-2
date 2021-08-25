package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lineNumber := 0
	for scanner.Scan() {
		txt := scanner.Text()
		lineNumber++
		fmt.Printf("%d : %s\n", lineNumber, txt)
	}
}
