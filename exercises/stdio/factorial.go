package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strconv"
)

func fact(n int64) int64 {
	if n == 1 {
		return 1
	}
	return n * fact(n - 1)
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var str string
	var number int64
	var err error

	fmt.Println("Entre com um número inteiro:")
	str, err = reader.ReadString(10)
	if err != nil {
		log.Fatal(err)
	}

	number, err = strconv.ParseInt(str[:len(str)-1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Fatorial de %v é %v.\n", number, fact(number))
}
