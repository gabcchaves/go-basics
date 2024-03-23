package main
import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "log"
)

func main() {
    var reader = bufio.NewReader(os.Stdin)
    var inputStr string
    var inputNum int64
    var err error
    
    fmt.Println("Entre com um número inteiro:")
    inputStr, _ = reader.ReadString('\n')
    inputNum, err = strconv.ParseInt(inputStr[:len(inputStr)-1], 10, 64)
    if err != nil {
        log.Fatal("Número inválido.")
    }
    
    if inputNum % 2 == 0 {
        if inputNum % 5 == 0 {
            fmt.Println("Termina com 0.")
        } else {
            fmt.Println("Não termina com 0.")
        }
    } else {
        if inputNum % 7 == 0 {
            fmt.Printf("Divisível apenas por 1, 7 e %v.", inputNum)
        } else {
            fmt.Println("Não é múltiplo de 7.")
        }
    }
}
