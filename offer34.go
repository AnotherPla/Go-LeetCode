/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 先序遍历 + 回溯
package main

func PathSum(root *TreeNode, target int) [][]int {
	var part func(*TreeNode, int)
	var res [][]int
	var row []int
	part = func(root *TreeNode, tar int) {
		if root == nil {
			return
		}
		row = append(row, root.Val)
		tar = tar - root.Val
		if tar == 0 && root.Left == nil && root.Right == nil {
			//tmp := make([]int,len(row))
			//copy(tmp,row)
			res = append(res, append([]int(nil), row...))
		}
		part(root.Left, tar)
		part(root.Right, tar)
		row = row[0 : len(row)-1]
	}
	part(root, target)
	return res
}
