package main

import (
	"fmt"
	circular "practice/linkedlist/Circular"
	doubly "practice/linkedlist/Doubly"
	singly "practice/linkedlist/Singly"
)

func main() {
	TestCircular()
}

func PrintSection(heading string) {
	fmt.Println("======================================================")
	fmt.Printf("%s\n", heading)
	fmt.Println("======================================================")
}

func TestCircular() {
	// ====================================
	// ============ Traversing ============
	// ====================================
	_, tailNode := circular.GetTestNodes()
	PrintSection("Traversing Circular Linkedlist (singly)")
	fmt.Printf("Traverse Forward: ")
	circular.Traverse(tailNode)
	fmt.Println()

	// ====================================
	// ============ Insertion =============
	// ====================================
	PrintSection("Insertion in circular LinkedList (singly)")

	fmt.Println("Inserting 10 in empty list")
	tail := circular.InsertIntoEmptyList(nil, 10)
	fmt.Printf("Traverse: ")
	circular.Traverse(tail)

	_, tailNode = circular.GetTestNodes()
	fmt.Println("Inserting 10 at beginning")
	tailNode = circular.InsertAtBeginning(tailNode, 10)
	fmt.Printf("Traverse: ")
	circular.Traverse(tailNode)

	_, tailNode = circular.GetTestNodes()
	fmt.Println("Inserting 100 at end")
	tailNode = circular.InsertAtEnd(tailNode, 100)
	fmt.Printf("Traverse: ")
	circular.Traverse(tailNode)

	_, tailNode = circular.GetTestNodes()
	fmt.Println("Inserting 50 at specific position 3 (index 2)")
	tailNode = circular.InsertAtPosition(tailNode, 3, 50)
	fmt.Printf("Traverse: ")
	circular.Traverse(tailNode)

	// ====================================
	// ============ Deletion ==============
	// ====================================
	PrintSection("Deletion in Circular LinkedList (singly)")

	_, tailNode = circular.GetTestNodes()
	fmt.Println("Deleting first node")
	tailNode = circular.DeleteFirstNode(tailNode)
	fmt.Printf("Traverse: ")
	circular.Traverse(tailNode)

	_, tailNode = circular.GetTestNodes()
	fmt.Println("Deleting last node")
	tailNode = circular.DeleteLastNode(tailNode)
	fmt.Printf("Traverse: ")
	circular.Traverse(tailNode)

	_, tailNode = circular.GetTestNodes()
	fmt.Println("Deleting node 4 (index 3)")
	tailNode = circular.DeleteAtPosition(tailNode, 3)
	fmt.Printf("Traverse Forward: ")
	circular.Traverse(tailNode)

	_, tailNode = circular.GetTestNodes()
	fmt.Println("Deleting node with key 3")
	tailNode = circular.DeleteKey(tailNode, 3)
	fmt.Printf("Traverse Forward: ")
	circular.Traverse(tailNode)
}

func TestDoubly() {
	// ====================================
	// ============ Traversing ============
	// ====================================
	headNode, tailNode := doubly.GetTestNodes()
	PrintSection("Traversing Doubly Linkedlist")
	fmt.Printf("Traverse Forward: ")
	doubly.TraverseForward(headNode)
	fmt.Printf("Traverse Forward Recursive: ")
	doubly.TraverseForwardRecursive(headNode)
	fmt.Println()
	fmt.Printf("Traverse Forward: ")
	doubly.TraverseBackward(tailNode)
	fmt.Printf("Traverse Backward Recursive: ")
	doubly.TraverseBackwardRecursive(tailNode)
	fmt.Println()

	// ====================================
	// ============== Length ==============
	// ====================================
	PrintSection("Finding length of Doubly LinkedList")
	length := doubly.Length(headNode)
	fmt.Printf("Length Iterative: %d\n", length)
	length = doubly.LengthRecursive(headNode)
	fmt.Printf("Length Iterative: %d\n", length)

	// ====================================
	// ============ Insertion =============
	// ====================================
	PrintSection("Insertion in doubly LinkedList")

	headNode, tailNode = doubly.GetTestNodes()
	fmt.Println("Inserting 10 at beginning")
	newHead := doubly.InsertAtBeginning(headNode, 10)
	fmt.Printf("Traverse Forward: ")
	doubly.TraverseForward(newHead)
	fmt.Printf("Traverse Backward: ")
	doubly.TraverseBackward(tailNode)

	headNode, _ = doubly.GetTestNodes()
	fmt.Println("Inserting 100 at end")
	newHead = doubly.InsertAtEnd(headNode, 100)
	fmt.Printf("Traverse Forward: ")
	doubly.TraverseForward(newHead)

	headNode, tailNode = doubly.GetTestNodes()
	fmt.Println("Inserting 50 at specific position 2 (index 1)")
	newHead = doubly.InsertAtPosition(headNode, 1, 50)
	fmt.Printf("Traverse Forward: ")
	doubly.TraverseForward(newHead)
	fmt.Printf("Traverse Backward: ")
	doubly.TraverseBackward(tailNode)

	// ====================================
	// ============ Deletion ==============
	// ====================================
	PrintSection("Deletion in Doubly LinkedList")

	headNode, tailNode = doubly.GetTestNodes()
	fmt.Println("Deleting first node")
	newHead = doubly.DeleteFirstNode(headNode)
	fmt.Printf("Traverse Forward: ")
	doubly.TraverseForward(newHead)
	fmt.Printf("Traverse Backward: ")
	doubly.TraverseBackward(tailNode)

	headNode, _ = doubly.GetTestNodes()
	fmt.Println("Deleting last node")
	newHead = doubly.DeleteLastNode(headNode)
	fmt.Printf("Traverse Forward: ")
	doubly.TraverseForward(newHead)

	headNode, tailNode = doubly.GetTestNodes()
	fmt.Println("Deleting node 4 (index 3)")
	newHead = doubly.DeleteAtPosition(headNode, 3)
	fmt.Printf("Traverse Forward: ")
	doubly.TraverseForward(newHead)
	fmt.Printf("Traverse Backward: ")
	doubly.TraverseBackward(tailNode)
}

func TestSingly() {
	nodeOne := singly.NewNode(1)
	nodeTwo := singly.NewNode(2)
	nodeThree := singly.NewNode(3)
	nodeFour := singly.NewNode(4)
	nodeFive := singly.NewNode(5)
	nodeSix := singly.NewNode(6)

	nodeOne.Next = nodeTwo
	nodeTwo.Next = nodeThree
	nodeThree.Next = nodeFour
	nodeFour.Next = nodeFive
	nodeFive.Next = nodeSix
	// ====================================
	// ============ Traversing ============
	// ====================================
	PrintSection("Traversing Singly Linkedlist")
	singly.Traverse(nodeOne)
	singly.TraverseRecursive(nodeOne)

	// ====================================
	// ============ Searching =============
	// ====================================
	PrintSection("Searching Singly LinkedList")
	key := 9
	fmt.Printf("Searching linkedlist for: %d\n", key)
	// result := singly.Search(nodeOne, key)
	result := singly.SearchRecursive(nodeOne, key)
	fmt.Printf("Found? %v\n", result)

	// ====================================
	// ============= Length ===============
	// ====================================
	PrintSection("Finding Singly LinkedList Length")
	// length := singly.Length(nodeOne)
	length := singly.LengthRecursive(nodeOne)
	fmt.Printf("Length of LinkedList: %d\n", length)

	// ====================================
	// ============ Insertion =============
	// ====================================
	PrintSection("Insertion in Singly LinkedList")
	fmt.Println("Inserting 10 at beginning")
	newHead := singly.InsertAtBeginning(nodeOne, 10)
	singly.Traverse(newHead)
	fmt.Println("Inserting 100 at end")
	newHead = singly.InsertAtEnd(nodeOne, 100)
	singly.Traverse(newHead)
	fmt.Println("Inserting 50 at specific position 2 (index 1)")
	newHead = singly.InsertAtPosition(nodeOne, 1, 50)
	singly.Traverse(newHead)

	// ====================================
	// ============ Deletion ==============
	// ====================================
	PrintSection("Deletion in Singly LinkedList")
	fmt.Println("Deleting first node")
	// newHead := singly.DeleteFirstNode(nodeOne)
	singly.Traverse(newHead)

	fmt.Println("Deleting last node")
	// newHead := singly.DeleteLastNode(nodeOne)
	singly.Traverse(newHead)

	fmt.Println("Deleting node 4 (index 3)")
	// newHead := singly.DeleteAtPosition(nodeOne, 5)
	if newHead != nil {
		singly.Traverse(newHead)
	} else {
		fmt.Println("Out of bounds")
	}
}
