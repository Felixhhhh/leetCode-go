package main

func main() {

	findDuplicates([]int{4,3,2,7,8,2,3,1})

}

func findDuplicates(nums []int) ( []int) {
	ans := make([]int,0)
	for _, x := range nums {
		if x < 0 {
			x = -x
		}
		if nums[x-1] > 0 {
			nums[x-1] = - nums[x-1]
		} else {
			ans = append(ans, x)
		}
	}
	return ans
}