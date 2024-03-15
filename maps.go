package main
import("fmt")

func main() {
	var a = map[int]string{23:"Peace", 11:"Chaos"}
	delete(a, 11)
	fmt.Println(a)
}
