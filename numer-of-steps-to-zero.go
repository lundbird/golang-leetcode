package main
import ( 
    "fmt"
) 
func main() {
    test := [][]int{
        {14},
        {8},
		{123},
		{0},
		{1},
        {2},
    }
    for i , _ := range test {
        fmt.Println(numberOfSteps(test[i][0]))
    }
}

func numberOfSteps(num int) int {
	if num==0 { return 0}
	ans := 1
	for q:=num; q>1; q=q/2 {
		ans += 1 + q%2
		// fmt.Println(q,ans)

	}
	return ans
}