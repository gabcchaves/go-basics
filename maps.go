package main
import("fmt")

func main() {
	var a = map[int]string{23:"Peace", 11:"Chaos"}
	for k, v := range a {
		fmt.Println(k, v)
	}
}
