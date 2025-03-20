package doubly

import "fmt"

type Node struct {
	Data int
	Next *Node
	Prev *Node
}

func NewNode(data int) *Node {
	return &Node{
		Data: data,
		Next: nil,
		Prev: nil,
	}
}

func GetTestNodes() (head, tail *Node) {
	nodeOne := NewNode(1)
	nodeTwo := NewNode(2)
	nodeThree := NewNode(3)
	nodeFour := NewNode(4)
	nodeFive := NewNode(5)
	nodeSix := NewNode(6)

	nodeOne.Prev = nil
	nodeOne.Next = nodeTwo

	nodeTwo.Prev = nodeOne
	nodeTwo.Next = nodeThree

	nodeThree.Prev = nodeTwo
	nodeThree.Next = nodeFour

	nodeFour.Prev = nodeThree
	nodeFour.Next = nodeFive

	nodeFive.Prev = nodeFour
	nodeFive.Next = nodeSix

	nodeSix.Prev = nodeFive
	nodeSix.Next = nil

	return nodeOne, nodeSix
}

func TraverseForward(head *Node) {
	if head == nil {
		fmt.Println("Nothing to traverse")
	}

	for current := head; current != nil; current = current.Next {
		fmt.Printf("%d ", current.Data)
	}
	fmt.Println("")
}

func TraverseForwardRecursive(head *Node) {
	if head == nil {
		return
	}
	fmt.Printf("%d ", head.Data)
	TraverseForwardRecursive(head.Next)
}

func TraverseBackward(tail *Node) {
	if tail == nil {
		fmt.Println("Nothing to traverse")
	}

	for current := tail; current != nil; current = current.Prev {
		fmt.Printf("%d ", current.Data)
	}
	fmt.Println("")
}

func TraverseBackwardRecursive(tail *Node) {
	if tail == nil {
		return
	}
	fmt.Printf("%d ", tail.Data)
	TraverseBackwardRecursive(tail.Prev)
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
		return nil
	}
	newNode := NewNode(value)

	head.Prev = newNode

	newNode.Next = head
	return newNode
}

func InsertAtEnd(head *Node, value int) *Node {
	newNode := NewNode(value)

	if head == nil {
		return nil
	}

	lastNode := &Node{}
	for current := head; current != nil; current = current.Next {
		lastNode = current
	}

	newNode.Prev = lastNode

	lastNode.Next = newNode
	return head
}

func InsertAtPosition(head *Node, pos, value int) *Node {
	newNode := NewNode(value)

	if head == nil {
		return nil
	}

	// If the position is less than 0, including negatives,
	// then add the new node as head.
	if pos < 1 {
		head.Prev = newNode
		newNode.Next = head
		return newNode
	}

	var lastNode *Node

	// If the position is greater than the length of the list, then add the node as tail.
	for i, current := 0, head; i < pos && current != nil; i++ {
		lastNode = current
		current = current.Next
	}

	nextNode := lastNode.Next
	if nextNode != nil {
		nextNode.Prev = newNode
	}
	newNode.Next = nextNode
	newNode.Prev = lastNode
	lastNode.Next = newNode

	return head
}

func DeleteFirstNode(head *Node) *Node {
	if head == nil {
		return nil
	}
	head = head.Next
	// Previous head no longer exist, remove reference to it.
	head.Prev = nil

	return head
}

func DeleteLastNode(head *Node) *Node {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return nil
	}

	var lastNode *Node
	// Loop until we are at second last Node
	for current := head; current != nil; current = current.Next {
		lastNode = current
	}
	previousNode := lastNode.Prev
	lastNode.Prev = nil
	previousNode.Next = nil

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
		newHead.Prev = nil
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

	// Node to be deleted
	toDeleteNode := lastNode.Next
	// Node after the node to be deleted
	nextNode := toDeleteNode.Next
	if nextNode != nil {
		nextNode.Prev = lastNode
	}
	// lastNode is the node before the to be deleted node
	lastNode.Next = nextNode
	toDeleteNode.Next = nil
	toDeleteNode.Prev = nil

	return head
}
