package main
import("fmt")

func main() {
	slice0 := []int{1, 2, 3}
	slice1 := []int{4, 5, 6}
	slice2 := append(slice0, slice1...)
	
	fmt.Println(slice2)
}
