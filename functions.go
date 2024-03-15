package main
import("fmt")

func sayHi(firstName string, lastName string) {
	fmt.Printf("Hi, %v %v.\n", firstName, lastName)
}

func main() {
	sayHi("Gabriel", "Chaves")
}
