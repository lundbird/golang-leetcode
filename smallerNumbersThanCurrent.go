package main
import ( 
	"fmt"
	"sort"
) 
func main() {
    test := [][]int{
		{8,1,2,2,3},
		{6,5,4,8},
		{7,7,7,7},
		{0,1},
    }
    for i , _ := range test {
        fmt.Println(smallerNumbersThanCurrent(test[i]))
    }
}

func smallerNumbersThanCurrent(nums []int) []int {
	orig := make([]int,len(nums))
	copy(orig,nums)
	sort.Ints(nums)

	ans := make([]int, len(nums))
	for i,_ := range orig {
		ans[i] = sort.SearchInts(nums, orig[i])
	}
	// fmt.Println(nums,orig,ans)
	return ans
}