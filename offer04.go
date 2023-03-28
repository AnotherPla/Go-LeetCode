package main

func FindNumberIn2DArray(matrix [][]int, target int) bool {
	var n, m int
	n = len(matrix)
	if n == 0 {
		return false
	} else {
		m = len(matrix[0])
	}
	var x, y = 0, m - 1
	for x < n && y >= 0 {
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] > target {
			y--
		} else {
			x++
		}
	}
	return false
}
