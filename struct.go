package main
import("fmt")

type Person struct {
	name string
	age int
}

func main() {
	var p Person
	p.name = "Gabriel"
	p.age = 22

	fmt.Printf("%v, %v.\n", p.name, p.age)
}
