package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fileContents, err := ioutil.ReadFile("data1.txt")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(fileContents))
}
