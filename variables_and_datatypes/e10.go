package main
import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    
    fmt.Println("Entre com um número inteiro:")
    input1, _ :=  reader.ReadString('\n')
    input1 = input1[:len(input1)-1]
    n1, err1 := strconv.ParseFloat(input1, 64)
    if err1 != nil {
        fmt.Println("Número inválido.")
    }
    
    fmt.Println("Entre com outro número inteiro:")
    input2, _ :=  reader.ReadString('\n')
    input2 = input2[:len(input2)-1]
    n2, err2 := strconv.ParseFloat(input2, 64)
    if err2 != nil {
        fmt.Println("Número inválido.")
    }

    fmt.Println("A soma dos números é igual a", n1 + n2)
} 
