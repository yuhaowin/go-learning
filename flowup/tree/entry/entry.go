package main

import (
	"fmt"
	"golearning/flowup/tree"
)

// MyTreeNode 通过组合的方式扩展已有的类型
type MyTreeNode struct {
	node *tree.Node
}

func (myNode *MyTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := MyTreeNode{myNode.node.Left}
	left.postOrder()

	right := MyTreeNode{myNode.node.Right}
	right.postOrder()

	myNode.node.Print()
}

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Print()

	root.Right.Left.SetValue(6)
	root.Right.Left.Print()

	root.SetValue(2)
	root.Print()

	nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)

	root.Traverse()

	fmt.Println()
	myRoot := MyTreeNode{&root}
	myRoot.postOrder()

	fmt.Println()
	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("Node Count:", nodeCount)

	fmt.Println("traverse with channel")

	maxValue := 0
	channel := root.TraverseWithChannel()
	for node := range channel {
		if node.Value > maxValue {
			maxValue = node.Value
		}
	}
	fmt.Println("Max node value:", maxValue)
}
