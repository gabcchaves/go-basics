package main
import("fmt")

func sum(a int, b int) (tag string, result int) {
	tag = "Sum is"
	result = a + b
	return
}

func main() {
	fmt.Println(sum(1, 2))
}
