package main
import "fmt"

func main() {
    var a int = 10
    var b int = 20
    var c int = 30

    if a > b && a > c {
        fmt.Printf("%v é maior do que %v e %v.", a, b, c)
    } else if b > a && b > c {
        fmt.Printf("%v é maior do que %v e %v.", b, a, c)
    } else if c > a && c > b {
        fmt.Printf("%v é maior do que %v e %v.", c, a, b)
    }
}
