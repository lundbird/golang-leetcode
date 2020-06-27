package main
import ( 
    "fmt"
) 
func main() {
	test := []int{2,3,5,1,3}
	fmt.Println(kidsWithCandies(test,3))
	test2 := []int{4,2,1,1,2}
	fmt.Println(kidsWithCandies(test2,1))
}

func max(x []int) int {
	max := x[0]
	for _,v := range x {
		if v > max {
			max = v
		}
	}
	return max
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandy := max(candies)
	ans := make([]bool,len(candies),len(candies))
	for i,_ := range candies {
		ans[i] = candies[i] + extraCandies >= maxCandy
	}
	return ans
}