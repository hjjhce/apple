package main

import "fmt"

func main() {

	list := []int{5, 9, 3, 1, 7, 8}
	fmt.Println(binaryTreeSort(list))
}

func bubbleSort() {

}

func selectionSort() {

}

func quickSort() {

}

func insertionSort() {
}

// 二叉树排序 begin
type tree struct {
	value       int
	left, right *tree
}

func binaryTreeSort(list []int) []int {
	var root *tree
	for _, v := range list {
		root = addBinaryTree(root, v)
		fmt.Printf("%v\n", root)
	}

	var sortSlice []int
	return appendSlice(sortSlice, root)
}

func addBinaryTree(t *tree, value int) *tree {
	if t == nil {
		t = &tree{value: value}
		return t
	}

	if value < t.value {
		t.left = addBinaryTree(t.left, value)
	} else {
		t.right = addBinaryTree(t.right, value)
	}

	return t
}

func appendSlice(s []int, t *tree) []int {
	if t != nil {
		s = appendSlice(s, t.left)
		s = append(s, t.value)
		s = appendSlice(s, t.right)
	}

	return s
}

// 二叉树排序 end
