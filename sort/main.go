package main

func main() {
	//ans := bubbleSort([]int{5, 7, 9, 2, 3})
	//ans := insertionSort1([]int{4, 1, 3, 1, 5, 2})
	ans := quickSort([]int{4, 1, 3, 1, 5, 2}, 0, 5)
	for _, an := range ans {
		println(an)
	}
}

/* 冒泡排序 */
func bubbleSort(nums []int) []int {
	// 外循环：当最大元素排序完成后，数组长度减一
	for i := len(nums) - 1; i > 0; i-- {
		flag := false // 初始化标志位
		// 内循环：冒泡操作
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				// 交换 nums[j] 与 nums[j + 1]
				nums[j], nums[j+1] = nums[j+1], nums[j] // 1. 将 nums[j] 向右移动一位
				flag = true
			}
		}
		if flag == false {
			break
		}
	}
	return nums
}

/* 插入排序 */
func insertionSort(nums []int) []int {
	// 外循环：待排序元素数量为 n-1, n-2, ..., 1
	for i := 1; i < len(nums); i++ {
		//获取为排序的元素base
		base := nums[i]
		j := i - 1
		//当nums[j]>base，左侧元素向右移动
		for j >= 0 && nums[j] > base {
			nums[j+1] = nums[j] // 1. 将 nums[j] 向右移动一位
			j--
		}
		nums[j+1] = base // 2. 将 base 赋值到正确位置
	}
	return nums
}

//快速排序
func quickSort(nums []int, left, right int) []int {
	// 子数组长度为 1 时终止
	for left < right {
		// 哨兵划分
		pivot := partition(nums, left, right)
		// 对两个子数组中较短的那个执行快排
		if pivot-left < right-right {
			partition(nums, left, pivot-1)
			left = pivot + 1
		} else {
			partition(nums, pivot+1, right)
			right = pivot - 1
		}
	}
	return nums
}

/* 哨兵划分 */
func partition(nums []int, left, right int) int {
	// 以 nums[left] 作为基准数
	i, j := left, right
	for i < j {
		for i < j && nums[j] >= nums[left] {
			j-- // 从右向左找首个小于基准数的元素
		}
		for i < j && nums[i] <= nums[left] {
			i++ // 从左向右找首个大于基准数的元素
		}
		// 元素交换
		nums[i], nums[j] = nums[j], nums[i]
	}
	// 将基准数交换至两子数组的分界线
	nums[i], nums[left] = nums[left], nums[i]
	return i // 返回基准数的索引
}
