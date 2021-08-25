package main

import (
	"io"
	"os"
)

type AlphaReader string

func (alphaReader AlphaReader) Read(p []byte) (int, error) {
	count := 0
	for idx := 0; idx < len(alphaReader); idx++ {
		if (alphaReader[idx] >= 'a' && alphaReader[idx] <= 'z') || (alphaReader[idx] >= 'A' && alphaReader[idx] <= 'Z') {
			p[count] = alphaReader[idx]
			count++
		}
	}
	return count, io.EOF
}
func main() {
	alphaReader := AlphaReader("This is a test string to test numbers like 123456 and special characters like &^%$#*&( are filtered")
	/*
		p := make([]byte, len(alphaReader))
		n, err := alphaReader.Read(p)
		if err == io.EOF {
			println(string(p[:n]))
			return
		}
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	*/
	io.Copy(os.Stdout, alphaReader)
}
