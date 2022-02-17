package trees

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	data int64

	left  *Node
	right *Node
}

type BinaryTree struct {
	Root *Node
}

func (bt *BinaryTree) Insert(data int64) *BinaryTree {
	if bt.Root == nil {
		bt.Root = &Node{data: data, left: nil, right: nil}
	} else {
		bt.Root.Insert(data)
	}
	return bt
}

func (n *Node) Insert(data int64) {

	if n == nil {
		return
	} else {
		if data <= n.data {
			if n.left == nil {
				n.left = &Node{data: data, left: nil, right: nil}
			} else {
				n.left.Insert(data)
			}
		} else {
			if n.right == nil {
				n.right = &Node{data: data, left: nil, right: nil}
			} else {
				n.right.Insert(data)
			}
		}
	}

}

func Search(root *Node, data int64) bool {
	if root == nil {
		return false
	}
	if root.data == data {
		return true
	}

	return Search(root.left, data) || Search(root.right, data)
}

func SearchNonRecursive(root *Node, data int64) bool {
	isPresent := false
	if root == nil {
		return isPresent
	}
	current := root

	for current != nil {
		if current.data == data {
			isPresent = true
			break
		}
		if current.data > data {
			current = current.left
		} else {
			current = current.right
		}
	}
	return isPresent
}

func PrintInOrder(node *Node) {

	if node != nil {
		PrintInOrder(node.left)
		fmt.Println(node.data)

		PrintInOrder(node.right)
	}
}

func PrintPostOrder(node *Node) {
	if node != nil {
		PrintPostOrder(node.left)
		PrintPostOrder(node.right)
		fmt.Println(node.data)
	}
}

func PrintPreOrder(node *Node) {
	if node != nil {
		fmt.Println(node.data)
		PrintPostOrder(node.left)
		PrintPostOrder(node.right)

	}
}

func (bt *BinaryTree) SetRoot(node *Node) {
	bt.Root = node
}

func (bt *BinaryTree) Delete(value int64) bool {
	if bt.Root == nil {
		return false
	}

	var parent *Node = nil
	var currentNode = bt.Root
	for currentNode != nil && currentNode.data != value {
		parent = currentNode
		if currentNode.data > value {
			currentNode = currentNode.left
		} else {
			currentNode = currentNode.right
		}
	}

	if currentNode == nil {
		return false
	} else if currentNode.left == nil && currentNode.right == nil {
		if bt.Root.data == currentNode.data {
			bt.SetRoot(nil)
			return true
		} else if currentNode.data < parent.data {
			parent.left = nil
			return true
		} else {
			parent.right = nil
			return true
		}
	} else if currentNode.right == nil {
		if bt.Root.data == currentNode.data {
			bt.SetRoot(currentNode.left)
			return true
		} else if currentNode.data < parent.data {
			parent.left = currentNode.left
			return true
		} else {
			parent.right = currentNode.left
			return true
		}
	} else if currentNode.left == nil {
		if bt.Root.data == currentNode.data {
			bt.SetRoot(currentNode.right)
			return true
		} else if currentNode.data < parent.data {
			parent.left = currentNode.right
			return true
		} else {
			parent.right = currentNode.right
			return true
		}
	} else {
		least := findtheLeastNode(currentNode.right)
		temp := least.data
		bt.Delete(temp)
		currentNode.data = temp
	}
	return true
}

func findtheLeastNode(currentNode *Node) *Node {
	temp := currentNode

	for temp.left != nil {
		temp = temp.left
	}
	return temp
}

func (bt *BinaryTree) MinIntheTree() *Node {

	current := bt.Root

	for current.left != nil {
		current = current.left
	}
	return current
}

func MinintheTreeRecursive(root *Node) *Node {

	if root == nil {
		return nil
	} else if root.left == nil {
		return root
	} else {
		return MinintheTreeRecursive(root.left)
	}
}

func KThMaxValue(root *Node, k int) int {
	kth := -1
	orderedData := InorderTraversal(root, "")
	orderedElements := strings.Split(orderedData, ",")
	if len(orderedElements)-k >= 0 {
		i, err := strconv.Atoi(orderedElements[len(orderedElements)-k])
		if err != nil {
			kth = i
		}

	}
	return kth
}

func InorderTraversal(root *Node, result string) string {

	if root.left != nil {
		InorderTraversal(root.left, result)
	}
	result = strconv.FormatInt(root.data, 10) + ","
	if root.right != nil {
		InorderTraversal(root.right, result)
	}

	return result
}
