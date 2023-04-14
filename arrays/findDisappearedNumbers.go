package main

func main() {

	findDisappearedNumbers([]int{1, 2, 2, 4,5,7,7})

}

func findDisappearedNumbers(nums []int) (ans []int) {
	n := len(nums)
	for _, v := range nums {
		v = (v - 1) % n //得到num值对应的下标
		nums[v] += n //num-1下标位置的数+n放入nums数组中
	}
	for i, v := range nums {
		if v <= n {
			//因为值是[1,n]而第一个for循环经过if判断后筛选得到
			//的是有问题的下标即从0开始但值从1，所以下面用i+1
			ans = append(ans, i+1)//i位置上的值i+1未曾出现过
		}
	}
	return
}