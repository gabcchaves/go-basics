package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var a, b int64
	var str1, str2 string
	
	// Ler primeiro número inteiro.
	fmt.Println("Entre com um inteiro:")
	str1, _ = reader.ReadString('\n')
	a, _ = strconv.ParseInt(str1[:len(str1)-1], 10, 64)
	
	// Ler segundo número inteiro.
	fmt.Println("Entre com outro inteiro:")
	str2, _ = reader.ReadString('\n')
	b, _ = strconv.ParseInt(str2[:len(str2)-1], 10, 64)

	// Exibir soma dos inteiros.
	fmt.Println("Soma:", a + b)
}
