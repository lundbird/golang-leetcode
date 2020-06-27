package main
import (
	"fmt"
	"strconv"
	"strings"
)
func main(){
	fmt.Println(subtractProductAndSum(234))
}

func subtractProductAndSum(n int) int {
	str_n := strconv.Itoa(n)
	str := strings.Split(str_n,"")
	for i:=0; i<len(str); i++ {
		fmt.Println(str[i])
	}
	prod, sum := 1, 0
	for _,v := range str {
		val, _ := strconv.Atoi(v)
		prod *= val
		sum += val
	}
	return prod - sum
}