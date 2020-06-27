package main
import ( 
    "fmt"
) 
func main() {
    test := [][]int{
        {5,0},
		{4,3},
		{1,7},
		{10,5},
		{1,0},
    }
    for i , _ := range test {
        fmt.Println(xorOperation(test[i][0],test[i][1]))
    }
}

func xorOperation(n int, start int) int {
	ans := start
	for i:=1; i<n; i++ {
		ans ^= (start + 2*i)
	}
	return ans
}