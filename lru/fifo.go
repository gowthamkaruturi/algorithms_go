package lru

import "fmt"

type fifo struct {
}

func (l *fifo) evict(c *Cache) interface{} {
	fmt.Println("Evicting by fifo strtegy")
	key := c.doublyLinkedList.Front().key
	c.doublyLinkedList.RemoveFromtheFront()
	return key
}

func (l *fifo) get(node *node, c *Cache) {
	fmt.Println("Shuffling doubly linked list due to get operation")
}

func (l *fifo) set(node *node, c *Cache) {
	fmt.Println("Shuffling doubly linked list due to set operation")
	c.doublyLinkedList.AddToEnd(node)
}

func (l *fifo) set_overwrite(node *node, value interface{}, c *Cache) {
	fmt.Println("Shuffling doubly linked list due to set_overwrite operation")
}
