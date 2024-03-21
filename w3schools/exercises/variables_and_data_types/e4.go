package main
import "fmt"

const PI = 3.14

func main() {
	var radius float32 = 3
	var area = PI * radius * radius

	fmt.Println(area)
}
