package main

// 递归做法
func VerifyPostorder(postorder []int) bool {
	var part func(int, int) bool
	part = func(left, right int) bool {
		if left >= right {
			return true
		}
		cur := left
		for cur < right && postorder[right] > postorder[cur] {
			cur++
		}
		mid := cur
		for cur < right && postorder[right] < postorder[cur] {
			cur++
		}
		return cur == right && part(left, mid-1) && part(mid, right-1)
	}
	return part(0, len(postorder)-1)
}
