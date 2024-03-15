package main
import("fmt")

func main() {
	slice := make([]int, 5, 10)
	slice = append(slice, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(slice)
}
