package main

import(
	"fmt"
	"time"
)

func main() {
	var currTime = time.Now()
	fmt.Printf("%T, %v\n", currTime, currTime)
}
