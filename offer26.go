package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 本题就是看B树深度遍历序列中有没有A树深度遍历序列的子序列
// 错误思路
/* func IsSubStructure(A *TreeNode, B *TreeNode) bool {
	seqA := DFS(A)
	seqB := DFS(B)
	flag := false
	if len(seqB) == 0 {
		return false
	}
	for i, val := range seqA {
		if val == seqB[0] {
			for j := 0; j < len(seqB); j++ {
				if seqB[j] != seqA[j+i] {
					break
				}
			}
			flag = true
			break
		}

	}

	return flag
} */

func DFS(root *TreeNode) []int {
	//map判断某个节点是否被访问过，便于回溯
	var isCounted = make(map[*TreeNode]bool)

	//深度遍历用栈实现
	var nodeStack []*TreeNode

	//遍历后的节点值
	var dfsSeq []int

	if root == nil {
		return dfsSeq
	}

	dfsSeq = append(dfsSeq, root.Val)
	isCounted[root] = true

	nodeStack = append(nodeStack, root)

	for len(nodeStack) != 0 {

		//栈顶元素
		cur := nodeStack[len(nodeStack)-1]

		//节点左右孩子都为空或者都被访问过则出栈
		if cur.Right == nil && cur.Left == nil {
			nodeStack = nodeStack[0 : len(nodeStack)-1]
			continue
		}
		if cur.Right != nil && cur.Left != nil && isCounted[cur.Right] && isCounted[cur.Left] {
			nodeStack = nodeStack[0 : len(nodeStack)-1]
			continue
		}
		if cur.Right != nil && cur.Left == nil && isCounted[cur.Right] {
			nodeStack = nodeStack[0 : len(nodeStack)-1]
			continue
		}
		if cur.Right == nil && cur.Left != nil && isCounted[cur.Left] {
			nodeStack = nodeStack[0 : len(nodeStack)-1]
			continue
		}

		//只有左孩子不为空且没有被访问过时才将左孩子入栈
		if cur.Left != nil && !isCounted[cur.Left] {
			isCounted[cur.Left] = true
			nodeStack = append(nodeStack, cur.Left)
			dfsSeq = append(dfsSeq, cur.Left.Val)
			continue
		}
		//只有左孩子为空或者被访问过，且右孩子不空时才将又孩子入栈
		if cur.Right != nil && (isCounted[cur.Left] || cur.Left == nil) {
			isCounted[cur.Right] = true
			nodeStack = append(nodeStack, cur.Right)
			dfsSeq = append(dfsSeq, cur.Right.Val)
		}

	}
	return dfsSeq
}
