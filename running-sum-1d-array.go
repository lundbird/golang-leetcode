package main
import ( 
    "fmt"
) 
func runningSum(nums []int) []int {
    for i, _ := range nums {
        if i == 0 {
            continue;
        }
        nums[i] = nums[i]+nums[i-1]
    }
    return nums
}

func main() {
    tests := [][]int{
        {1,2,3,4},
        {1,1,1,1},
        {1},
    }
    for i , _ := range tests {
        fmt.Println(runningSum(tests[i]))
    }
}
