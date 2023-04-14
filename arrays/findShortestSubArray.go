package main

import (
	"math"
)

func main() {

	findShortestSubArray([]int{1,2,2,4})
}


func findShortestSubArray(nums []int) int {
	con := map[int]int{}
	ans := []int{math.MaxInt64, math.MaxInt64}
	for _, v := range nums {
		con[v]++
	}
	for k, v := range con {
		if v != 1 && v < ans[1] {
			ans[0], ans[1] = k, v
		}
	}
	length := 0
	low := 0
	hight := 0
	for i := 0; i < len(nums); i++ {
		if ans[0] == nums[i] && length == 0 {
			low = i
		}
		if ans[0] == nums[i] {
			length++
		}
		if length == ans[1] {
			hight = i
			break
		}
	}
	return hight - low + 1
}