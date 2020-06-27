package main
import ( 
    "fmt"
) 
func shuffle(nums []int, n int) []int {
    x := nums[:n]
	y := nums[n:]
	z := make([]int, n*2)
	for i:=0; i<n*2; i++ {
		if i % 2 == 0{
			z[i] = x[i/2]
		}else{
			z[i] = y[i/2]
		}
	}
	return z
}

func main() {
    test := [][]int{
        {1,2,3,4},
        {2,5,1,3,4,7},
		{1,2},
		{1,2,3,4,4,3,2,1},
		{1,1,2,2},
    }
    for i , _ := range test {
        fmt.Println(shuffle(test[i],len(test[i])/2))
    }
}