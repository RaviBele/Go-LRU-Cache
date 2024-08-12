package main

import (
	"fmt"
)

const SIZE = 5

type Node struct {
	value string
	left  *Node
	right *Node
}

type Queue struct {
	head   *Node
	tail   *Node
	length int
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}
	head.right = tail
	tail.left = head

	return Queue{head: head, tail: tail}
}

type Cache struct {
	queue Queue
	hash  HashMap
}

func NewCache() Cache {
	return Cache{queue: NewQueue(), hash: HashMap{}}
}

func (c *Cache) Check(key string) {
	node := &Node{}
	if val, ok := c.hash[key]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{value: key}
	}
	c.Add(node)
	c.hash[key] = node
}

func (c *Cache) Remove(node *Node) *Node {
	fmt.Printf("Remove %s\n", node.value)
	left := node.left
	right := node.right

	left.right = right
	right.left = left

	c.queue.length -= 1
	delete(c.hash, node.value)
	return node
}

func (c *Cache) Add(node *Node) {
	fmt.Printf("Add %s\n", node.value)
	headRight := c.queue.head.right

	node.left = c.queue.head
	node.right = headRight
	headRight.left = node
	c.queue.head.right = node
	c.queue.length += 1

	if c.queue.length > SIZE {
		c.Remove(c.queue.tail.left)
	}
}

func (c *Cache) Display() {
	node := c.queue.head.right

	for i := 0; i < c.queue.length; i++ {
		fmt.Printf("{%s}", node.value)
		if i < c.queue.length-1 {
			fmt.Printf("<-->")
		}
		node = node.right
	}
	fmt.Println()
}

type HashMap map[string]*Node

func main() {
	fmt.Println("Start LRU Cache with size: ", SIZE)
	cache := NewCache()

	for _, word := range []string{"ravi", "shruti", "kalpna", "mukta", "rohan", "varsha"} {
		cache.Check(word)
		cache.Display()
	}
}
