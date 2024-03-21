package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strconv"
)

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

	if number % 2 == 0 {
		fmt.Println("Número Par.")
	} else {
		fmt.Println("Número Ímpar.")
	}
}
