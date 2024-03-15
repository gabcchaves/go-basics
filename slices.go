package main
import("fmt")

func main() {
	var array = [...]int{1, 2, 3, 4, 5}
	slice := array[0:3]
	fmt.Println(slice)
}
