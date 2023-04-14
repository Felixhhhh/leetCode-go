package main

import (
	"fmt"
)

func main() {

	firstMissingPositive([]int{1, 2, 0, 5, 4})
}

func firstMissingPositive(nums []int) int {

	n := len(nums)
	//若存在1 直接返回
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}
	//将小于1或大于n的数据置为1

	for i := 0; i < n; i++ {
		num := abs(nums[i])
		if num <= n {
			fmt.Println(num - 1)
			nums[num-1] = -abs(nums[num-1])
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
