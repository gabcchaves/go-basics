package main
import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "log"
)

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
    var inputStr string
    var inputRune rune
    var a int64
    var b int64
    var err error
    var repeat bool = true

    for repeat {
        fmt.Println("Qual operação? [+, -, *, /, s(sair)]")
        inputRune, _, err = reader.ReadRune()
        if err != nil {
            log.Fatal("Erro de entrada.")
        }
        _, _, err = reader.ReadRune()
        if inputRune == 's' {
            os.Exit(0)
        }

        fmt.Println("Entre com um número inteiro:")
        inputStr, err  = reader.ReadString('\n')
        if err != nil {
            log.Fatal("Erro de entrada.")
        }
        a, err = strconv.ParseInt(inputStr[:len(inputStr)-1], 10, 64)
        if err != nil {
            log.Fatal("Número inválido.")
        }

        fmt.Println("Entre com outro número inteiro:")
        inputStr, err = reader.ReadString('\n')
        if err != nil {
            log.Fatal("Erro de entrada.")
        }
        b, err = strconv.ParseInt(inputStr[:len(inputStr)-1], 10, 64)
        if err != nil {
            log.Fatal("Número inválido.")
        }

        switch inputRune {
        case '+':
            fmt.Println("A soma é igual a", sum(a, b))
        case '-':
            fmt.Println("A subtração é igual a", subtract(a, b))
        case '*':
            fmt.Println("A multiplicação é igual a", multiply(a, b))
        case '/':
            fmt.Println("A divisão é igual a", divide(a, b))
        default:
            fmt.Println("Operação inválida.")
        }
    }
}
