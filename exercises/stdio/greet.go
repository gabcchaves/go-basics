package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var name string
	var err error

	name, err = reader.ReadString(10)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Olá, %v.\n", name)
}
