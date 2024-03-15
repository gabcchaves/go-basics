package main
import("fmt")

func main() {
	fmt.Println("HI");
	for i:=0; i < 5; i++ {
		fmt.Println(i)
	}

	for i:=5; i > 0; i-- {
		fmt.Println(i)
	}

	for i:=0; i < 5; i++ {
		fmt.Println("For loop");
	}
}
