package singly

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func NewNode(data int) *Node {
	return &Node{
		Data: data,
		Next: nil,
	}
}

func Traverse(head *Node) {
	if head == nil {
		fmt.Println("Nothing to traverse")
	}

	for current := head; current != nil; current = current.Next {
		fmt.Printf("%d ", current.Data)
	}
	fmt.Println("")
}

func TraverseRecursive(head *Node) {
	if head == nil {
		return
	}
	fmt.Printf("%d ", head.Data)
	TraverseRecursive(head.Next)
}

func Search(head *Node, key int) bool {
	if head == nil {
		fmt.Println("Nothing to search")
		return false
	}

	for current := head; current != nil; current = current.Next {
		if current.Data == key {
			return true
		}
	}

	return false
}

func SearchRecursive(head *Node, key int) bool {
	if head == nil {
		return false
	}

	if head.Data == key {
		return true
	}

	return SearchRecursive(head.Next, key)
}

func Length(head *Node) int {
	count := 0
	if head == nil {
		return count
	}

	for current := head; current != nil; current = current.Next {
		count++
	}
	return count
}

func LengthRecursive(head *Node) int {
	if head == nil {
		return 0
	}

	return 1 + LengthRecursive(head.Next)
}

func InsertAtBeginning(head *Node, value int) *Node {
	if head == nil {
		return &Node{}
	}
	newNode := NewNode(value)
	newNode.Next = head
	return newNode
}

func InsertAtEnd(head *Node, value int) *Node {
	newNode := NewNode(value)

	if head == nil {
		return newNode
	}

	lastNode := &Node{}
	for current := head; current != nil; current = current.Next {
		lastNode = current
	}

	lastNode.Next = newNode
	return head
}

func InsertAtPosition(head *Node, pos, value int) *Node {
	newNode := NewNode(value)

	if head == nil {
		return newNode
	}

	// If the position is less than 0, including negatives,
	// then add the new node as head.
	if pos < 1 {
		newNode.Next = head
		return newNode
	}

	var lastNode *Node

	// If the position is greater than the length of the list, then add the node as tail.
	for i, current := 0, head; i < pos && current != nil; i++ {
		lastNode = current
		current = current.Next
	}

	newNode.Next = lastNode.Next
	lastNode.Next = newNode

	return head
}

func DeleteFirstNode(head *Node) *Node {
	if head == nil {
		return nil
	}
	head = head.Next
	return head
}

func DeleteLastNode(head *Node) *Node {
	if head == nil {
		return nil
	}

	var secondLastNode *Node
	// Loop until we are at second last Node
	for current := head; current.Next != nil; current = current.Next {
		secondLastNode = current
	}
	secondLastNode.Next = nil

	return head
}

func DeleteAtPosition(head *Node, pos int) *Node {
	if head == nil {
		return nil
	}

	// If the position is less than 0, including negatives,
	// then remove the first node
	// Alternative: Return an error
	if pos < 1 {
		newHead := head.Next
		head.Next = nil
		return newHead
	}

	var lastNode *Node

	// If the position is greater than the length of the list, then add the node as tail.
	for i, current := 0, head; i < pos && current != nil; i++ {
		lastNode = current
		current = current.Next
	}

	if lastNode.Next == nil {
		return nil
	}

	toDeleteNode := lastNode.Next
	lastNode.Next = toDeleteNode.Next
	toDeleteNode.Next = nil

	return head
}
