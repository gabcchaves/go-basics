package main
import("fmt")

func main() {
	slice0 := []int{1, 2, 3}
	
	fmt.Println(slice0)

	slice0 = slice0[0:2]

	fmt.Println(slice0)
}
