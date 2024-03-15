package main
import("fmt")

func sum(a int, b int) (result int) {
	result = a + b
	return
}

func main() {
	fmt.Println(sum(1, 2))
}
