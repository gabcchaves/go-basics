package main
import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strconv"
)

func ReadInt(reader *bufio.Reader) int64 {
	var inputStr string
	var a int64
	var err error
	inputStr, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal("Erro na leitura do número.")
	}
	a, err = strconv.ParseInt(inputStr[:len(inputStr)-1], 10, 64)
	if err != nil {
		log.Fatal("Erro na conversão de string para inteiro.")
	}
	return a
}

func ReadRune(reader *bufio.Reader) rune {
	inputRune, _, err := reader.ReadRune()
	if err != nil {
		log.Fatal("Erro na leitura da opção.")
	}
	_, _, _ = reader.ReadRune()
	return inputRune
}

func sum(a int64, b int64) int64 {
    return a + b
}

func subtract(a int64 , b int64) int64 {
    return a - b
}

func multiply(a int64, b int64 ) int64 {
    return a * b
}

func divide(a int64, b int64) int64 {
    return a / b
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var inputRune rune
	var a, b int64
	var exit bool = false

	fmt.Println("Calculadora")

	for !exit {
		fmt.Println("Escolha uma operação: +, -, *, /, s")
		inputRune = ReadRune(reader)

		switch inputRune {
		case '+':
			fmt.Println("Entre com um número inteiro.")
			a = ReadInt(reader)
			fmt.Println("Entre com outro número inteiro.")
			b = ReadInt(reader)
			fmt.Printf("A soma de ambos é igual a %v\n", sum(a, b))
		case '-':
			fmt.Println("Entre com um número inteiro.")
			a = ReadInt(reader)
			fmt.Println("Entre com outro número inteiro.")
			b = ReadInt(reader)
			fmt.Printf("A soma de ambos é igual a %v\n", subtract(a, b))
		case '*':
			fmt.Println("Entre com um número inteiro.")
			a = ReadInt(reader)
			fmt.Println("Entre com outro número inteiro.")
			b = ReadInt(reader)
			fmt.Printf("A soma de ambos é igual a %v\n", multiply(a, b))
		case '/':
			fmt.Println("Entre com um número inteiro.")
			a = ReadInt(reader)
			fmt.Println("Entre com outro número inteiro.")
			b = ReadInt(reader)
			fmt.Printf("A soma de ambos é igual a %v\n", divide(a, b))
		case 's':
			exit = true
		default:
			fmt.Println("Operação inválida.")
		}
	}
}
