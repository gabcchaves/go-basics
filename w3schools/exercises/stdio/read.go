package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	fmt.Println("Enter a string:")
	var reader = bufio.NewReader(os.Stdin)
	var str, err = reader.ReadString(10)
	fmt.Println(str, err)
}
