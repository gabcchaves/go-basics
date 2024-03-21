package main

import (
	"fmt"
	"math"
)

const PI = 3.1415

func main() {
	var radius float64 = 3
	var sphereVolume float64 = 4 * PI * math.Pow(radius, 3) / 3
	fmt.Println(sphereVolume)
}
