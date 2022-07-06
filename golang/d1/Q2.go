package main

import "fmt"

type Node struct {
	data        string
	left, right *Node
}

func createNewNode(data string, left, right *Node) *Node {
	node := new(Node)
	node.data = data
	node.left = left
	node.right = right
	return node
}

func preorder(currNode *Node, res *string) {
	if currNode == nil {
		return
	}
	*res += currNode.data + " "
	preorder(currNode.left, res)
	preorder(currNode.right, res)
}

func postorder(currNode *Node) {
	if currNode == nil {
		return
	}
	postorder(currNode.left)
	postorder(currNode.right)
	fmt.Print(currNode.data, " ")
}

func main() {

	leftTree := createNewNode("a", nil, nil)
	rightTree := createNewNode("-", createNewNode("b", nil, nil), createNewNode("c", nil, nil))
	root := createNewNode("+", leftTree, rightTree)

	res := ""
	preorder(root, &res)
	fmt.Println()
	postorder(root)
}
