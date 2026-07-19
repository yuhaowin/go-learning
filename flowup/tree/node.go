package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

// 为结构体定义方法，这里的 node 为接受者
// 这里是值接受者
func (node treeNode) print() {
	fmt.Println(node.value)
}

// 这里是指针接受者
func (node *treeNode) setValue(value int) {
	node.value = value
}

func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

func main() {
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)
	root.print()

	root.right.left.setValue(5)
	root.right.left.print()

	root.setValue(2)
	root.print()

	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)
}
