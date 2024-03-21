package main
import(
	"fmt"
	"strconv"
)

func main() {
	var str string = "10"
	var integer, err = strconv.ParseInt(str, 10, 64)
	if err == nil {
		fmt.Println(integer)
	}
}
