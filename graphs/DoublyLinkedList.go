package graphs

import (
	"fmt"
	"strconv"
)

type DoublyLinkedList struct {
	tail *Node
	head *Node
	size int
}

func initDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (d *DoublyLinkedList) AddFromFront(data int) {
	newnode := &Node{data: data}

	if d.head == nil {
		d.head = newnode
		d.tail = newnode
	} else {
		newnode.next = d.head
		d.head.prev = newnode
		d.head = newnode
	}
	d.size++
	return
}

func (d *DoublyLinkedList) AddFromEnd(data int) {
	newnode := &Node{data: data}

	if d.head == nil {
		d.head = newnode
		d.tail = newnode
	} else {
		currentNode := d.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		newnode.prev = currentNode
		currentNode.next = newnode
		d.tail = newnode
	}
	d.size++

}

func (d *DoublyLinkedList) GetHeadNode() *Node {
	return d.head
}

func (d *DoublyLinkedList) GetTail() *Node {
	return d.tail
}

func (d *DoublyLinkedList) IsEmpty() bool {
	isEmpty := false
	if d.tail == nil && d.head == nil {
		isEmpty = true
	}
	return isEmpty
}
func (d *DoublyLinkedList) Print() {
	if d.IsEmpty() {
		fmt.Println("there is no data in graph")
	} else {
		current := d.head

		for current != nil {
			fmt.Print(strconv.Itoa(current.data) + "<-->")
			current = current.next

		}
		fmt.Println("null")
	}
}

func (d *DoublyLinkedList) PrintReverse() {
	if d.IsEmpty() {
		fmt.Println("there is no data in graph")
	} else {
		current := d.tail

		for current.next != nil {
			fmt.Print(strconv.Itoa(current.data) + "<-->")
			current = current.prev

		}
	}
}

type Node struct {
	data int
	prev *Node
	next *Node
}
