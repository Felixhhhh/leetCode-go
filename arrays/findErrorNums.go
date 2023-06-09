package main

import (
	"sort"
)

func main() {

	findErrorNums([]int{1,2,2,4})

}


func findErrorNums(nums []int) []int {
	ans := make([]int, 2)
	sort.Ints(nums)
	pre := 0
	for _, v := range nums {
		if v == pre {
			ans[0] = v
		} else if v-pre > 1 {
			ans[1] = pre + 1
		}
		pre = v
	}
	n := len(nums)
	if nums[n-1] != n {
		ans[1] = n
	}
	return ans
}