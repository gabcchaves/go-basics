package main
import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var inputChar rune
	var reader = bufio.NewReader(os.Stdin)
	
	fmt.Println("Entre com uma opção: (s, p)")
	inputChar, _, _ = reader.ReadRune()
	_ = reader.UnreadRune()

	switch inputChar {
	case 's':
		os.Exit(0)
	case 'p':
		fmt.Println("Imprimindo algo.")
	default:
		os.Exit(0)
	}
}
