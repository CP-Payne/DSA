// Circular contains functions that perform oprations on a circular (singlY) linkedlist
package circular

import (
	"fmt"
	singly "practice/linkedlist/Singly"
)

// GetTestNodes returns a circular linkedlist (singly)
func GetTestNodes() (head, tail *singly.Node) {
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
	// Linking last node with first
	nodeSix.Next = nodeOne

	return nodeOne, nodeSix
}

func Traverse(last *singly.Node) {
	if last == nil {
		fmt.Println("Nothing to traverse")
		return
	}

	headNode := last.Next

	for current := headNode; ; current = current.Next {

		fmt.Printf("%d ", current.Data)
		if current.Next == headNode {
			break
		}
	}
	fmt.Println()
}

func InsertIntoEmptyList(last *singly.Node, data int) (lastNode *singly.Node) {
	// If list is not nil, return the list without modification
	if last != nil {
		return last
	}

	newNode := singly.NewNode(data)

	newNode.Next = newNode

	last = newNode

	return last
}

func InsertAtBeginning(last *singly.Node, value int) *singly.Node {
	if last == nil {
		return nil
	}
	newNode := singly.NewNode(value)
	newNode.Next = last.Next
	last.Next = newNode
	return last
}

func InsertAtEnd(last *singly.Node, value int) *singly.Node {
	if last == nil {
		return nil
	}
	newNode := singly.NewNode(value)
	newNode.Next = last.Next
	last.Next = newNode
	last = newNode
	return last
}

func InsertAtPosition(last *singly.Node, pos, value int) *singly.Node {
	if last == nil {
		return nil
	}

	newNode := singly.NewNode(value)

	if pos < 0 {
		fmt.Println("Invalid Position")
		return last
	}

	// If the position is less than 0, including negatives,
	// then add the new node as head.
	if pos == 0 {
		newNode.Next = last.Next
		last.Next = newNode
		return last
	}

	headNode := last.Next

	current := headNode

	// If the position is greater than the length of the list, then add the node as tail.
	for i := 0; i < pos-1; i++ {
		current = current.Next

		if current == headNode {
			fmt.Println("Invalid Position")
			return last
			// break
		}
	}

	// if the index is greater than the length of the list add the node to the end
	// if lastTracked.Next == headNode {
	// 	lastTracked.Next = newNode
	// 	newNode.Next = headNode
	// 	last = newNode
	// 	return last
	// }
	//
	// If inserted at end
	if current == last {
		last.Next = newNode
		newNode.Next = headNode
		last = newNode
		return last
	}

	newNode.Next = current.Next
	current.Next = newNode

	return last
}

func DeleteFirstNode(last *singly.Node) *singly.Node {
	if last == nil {
		return nil
	}
	if last.Next == last {
		last = nil
		return last
	}

	toDelete := last.Next
	last.Next = toDelete.Next
	toDelete.Next = nil

	return last
}

func DeleteLastNode(last *singly.Node) *singly.Node {
	if last == nil {
		return nil
	}
	if last.Next == last {
		last = nil
		return last
	}

	// Start with head
	secondLastNode := last.Next
	for current := last.Next; current != last; current = current.Next {
		secondLastNode = current
	}

	secondLastNode.Next = last.Next
	last = secondLastNode
	return last
}

func DeleteAtPosition(last *singly.Node, pos int) *singly.Node {
	if last == nil {
		return nil
	}

	if last == last.Next {
		last = nil
		return last
	}

	if pos < 0 {
		fmt.Println("Invalid position")
		return last
	}

	current := last.Next

	// If the position is greater than the length of the list, then add the node as tail.
	for i := 0; i < pos-1; i++ {
		current = current.Next
		if current == last {
			fmt.Println("Invalid Position")
			return last
		}
	}

	toDelete := current.Next
	current.Next = toDelete.Next
	if toDelete == last {
		last = current
	}

	toDelete.Next = nil

	return last
}

func DeleteKey(last *singly.Node, key int) *singly.Node {
	if last == nil {
		return nil
	}

	current := last.Next
	prev := last

	//
	if current == last && current.Data == key {
		last = nil
		return last
	}

	// If first node
	if current.Data == key {
		last.Next = current.Next
		return last
	}

	for current != last && current.Data != key {
		prev = current
		current = current.Next
	}

	if current.Data == key {
		prev.Next = current.Next
		if current == last {
			last = prev
		}
	} else {
		fmt.Println("Key not found")
	}
	return last
}
