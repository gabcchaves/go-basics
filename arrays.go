package main

import("fmt")

func main() {
	var a = [3]int{1, 2, 3}
	var b = [...]int{1, 2, 3, 4, 5}

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(a[0])
	fmt.Println(b[1])

	a[2] = 4

	fmt.Println(a[2])

	fmt.Println()

	var c = [5]int{}
	var d = [5]int{1, 2}
	fmt.Println(c)
	fmt.Println(d)
}
