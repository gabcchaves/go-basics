package main
import("fmt")

func sayHi(name string) {
	fmt.Printf("Hi, %v.\n", name)
}

func main() {
	sayHi("Gabriel")
}
