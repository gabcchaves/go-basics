package main
import("fmt")

func main() {
	slice0 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sliceCopy := make([]int, 5)
	copy(sliceCopy, slice0[0:6])
	fmt.Println(sliceCopy)
}
