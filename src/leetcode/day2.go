package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func conectRightAndLeft(node1 *Node, node2 *Node) {
	if node1 != nil {
		if node1.Right != nil {
			node1.Right.Next = node2.Left
			conectRightAndLeft(node1.Right, node2.Left)
		}
	}
}

func connectLeftAndright(node *Node) {
	if node.Left != nil {
		node.Left.Next = node.Right
		connectLeftAndright(node.Left)
		connectLeftAndright(node.Right)
	}

}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	connectLeftAndright(root)
	conectRightAndLeft(root.Left, root.Right)
	return root
}
