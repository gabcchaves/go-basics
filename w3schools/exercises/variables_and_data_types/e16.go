package main

import "fmt"

func main() {
	var a bool = true
	var b bool = true

	fmt.Println(a && b)
	fmt.Println(a || b)
	fmt.Println(a && !b)
	fmt.Println(a || !b)
}
