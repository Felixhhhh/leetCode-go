package leetcode

func generateMatrix(n int) [][]int {
	upBound, leftBound := 0, 0
	lowerBound, rightBound := n-1, n-1
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	num := 1
	for num <= n*n {
		if upBound <= lowerBound {
			for i := leftBound; i <= rightBound; i++ {
				res[upBound][i] = num
				num++
			}
			upBound++
		}
		if leftBound <= rightBound {
			for i := upBound; i <= lowerBound; i++ {
				res[i][rightBound] = num
				num++
			}
			rightBound--
		}
		if upBound <= lowerBound {
			for i := rightBound; i >= leftBound; i-- {
				res[lowerBound][i] = num
				num++
			}
			lowerBound--
		}
		if leftBound <= rightBound {
			for i := lowerBound; i >= upBound; i-- {
				res[i][leftBound] = num
				num++
			}
			leftBound++
		}
	}
	return res

}
