package main
import "fmt"

func main() {
	var i string = "10"
	var f, err = float64(i)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Erro de conversão.")
		}
	}()

	fmt.Println(f + 2.2)
	panic("Something went wrong.")
}
