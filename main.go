package main

import (
	"fmt"
)

func main() {
	fmt.Print("running...............................\n")
	//   05
	// ReplaceSpace("We are happy")

	//   03
	// FindRepeatNumber()

	//   58
	/* s := ReverseLeftWords("adawfgesg", 3)
	fmt.Printf("s: %s\n", s) */

	/* a := [][]int{{1, 2, 3}, {2, 3, 4}}
	fmt.Printf("len(a): %v\n", len(a))
	fmt.Printf("a[0]: %v\n", len(a[0])) */

	//   04  Z字形查找
	// FindNumberIn2DArray()

	//   11  二分
	// MinArray()

	//   50
	/* var s = "uindrseqbljlhqvlwvgdebeihttirikuahlikgnahvgnptmqwbovmuwesxkvcitcwrwrucsbbfqvldridfviduqvmfcmeiphoq
	upbitnwdbvevouoaetisdmgvvvwoglwtgjrpcbghxkrkjthetxeexbphbjiehaicuicgnirslhdstgmqcdnlulpdpadjdltfouwhfqicfcqntnpeq
	aohslwmeqmtdhnbkwnidigxdantmkckijiecavkpumeg"
	fmt.Printf("i: %c\n", i) */

	/* a := make(map[string]bool)

	k, v := a["tom"]
	fmt.Printf("k: %v\n", k)
	fmt.Printf("v: %v\n", v)
	fmt.Printf("a[\"tom\"]: %v\n", a["tom"])

	if !a["tom"] {
		fmt.Print("...")
	} */

	//DFS遍历二叉树
	/* a := TreeNode{Val: 1}
	b := TreeNode{Val: 0}
	c := TreeNode{Val: 1}
	d := TreeNode{Val: -4}
	e := TreeNode{Val: -3}

	a.Left = &b
	a.Right = &c
	b.Left = &d
	b.Right = &e

	res := DFS(&a)
	fmt.Printf("res: %v\n", res)

	f := TreeNode{Val: 1}
	g := TreeNode{Val: -4}
	f.Left = &g
	res2 := DFS(&f)
	fmt.Printf("res2: %v\n", res2) */

	// 20
	//IsNumber(".")

	/* s := strings.Split("a..a", ".")
	for _, v := range s {
		fmt.Printf("v:%v\n", v)
	}
	*/

	// 30
	/* post := []int{4, 8, 6, 12, 16, 14, 10}
	res := VerifyPostorder(post)
	fmt.Printf("res: %v\n", res) */
}
