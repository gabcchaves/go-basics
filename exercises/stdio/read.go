package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	fmt.Println("Enter a string:")
	var reader = bufio.NewReader(os.Stdin)
	var str, a, b = reader.ReadLine()
	fmt.Println(string(str), a, b)
}
