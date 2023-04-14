package main

// 差分数组工具类
type Difference struct {
	// 差分数组
	diff []int
}

/* 输入一个初始数组，区间操作将在这个数组上进行 */
func NewDifference(nums []int) *Difference {
	diff := make([]int, len(nums))
	// 根据初始数组构造差分数组
	diff[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		diff[i] = nums[i] - nums[i-1]
	}
	return &Difference{diff: diff}
}

/* 给闭区间 [i, j] 增加 val（可以是负数）*/
func (d *Difference) Increment(i, j, val int) {
	d.diff[i] += val
	if j+1 < len(d.diff) {
		d.diff[j+1] -= val
	}
}

func (d *Difference) Result() []int {
	res := make([]int, len(d.diff))
	res[0] = d.diff[0]
	// 根据差分数组构造结果数组
	for i := 1; i < len(res); i++ {
		res[i] = res[i-1] + d.diff[i]
	}
	return res
}
