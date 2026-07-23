package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

// 为结构体定义方法，这里的 node 为接受者
// 这里是值接受者
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

// 这里是指针接受者
func (node *Node) SetValue(value int) {
	node.Value = value
}

// 工厂方法
func CreateNode(value int) *Node {
	return &Node{Value: value}
}
