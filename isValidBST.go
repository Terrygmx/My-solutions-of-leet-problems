package main

//import (
//	"fmt"
//)

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	l := []int{}
	if root == nil {
		return l
	}
	if root.Left != nil {
		l = append(l, inorderTraversal(root.Left)...)
	}
	l = append(l, root.Val)
	if root.Right != nil {
		l = append(l, inorderTraversal(root.Right)...)
	}
	return l
}
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	l := inorderTraversal(root)
	for i := 0; i < len(l)-1; i++ {
		if l[i+1] < l[i] {
			return false
		}
	}
	return true
}

//func main() {
//	t1 := TreeNode{Val: 1}
//	t2 := TreeNode{Val: 3}
//	t3 := TreeNode{Val: 2, Left: &t1, Right: &t2}
//	fmt.Println(isValidBST(&t3))
//}
