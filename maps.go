package main
import("fmt")

func main() {
	var a = map[int]string{23:"Peace", 11:"Chaos"}
	val, ok := a[23]
	fmt.Println(val, ok)
}
