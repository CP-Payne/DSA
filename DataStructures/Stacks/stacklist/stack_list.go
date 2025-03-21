package stacklist

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
		Next:  nil,
	}
}

type Stack struct {
	Head *Node
}

func NewStack() *Stack {
	return &Stack{
		Head: nil,
	}
}

func (s *Stack) Push(value int) {
	newNode := NewNode(value)

	if s.Head == nil {
		s.Head = newNode
		return
	}
	newNode.Next = s.Head
	s.Head = newNode
}

func (s *Stack) Pop() int {
	if s.Head == nil {
		fmt.Println("Stack is empty")
		return 0
	}

	prevHead := s.Head
	s.Head = s.Head.Next
	prevHead.Next = nil

	return prevHead.Value
}

func (s *Stack) Peek() int {
	if s.Head == nil {
		fmt.Println("Stack is empty")
		return 0
	}

	return s.Head.Value
}

func (s *Stack) Display() {
	if s.Head == nil {
		fmt.Println("Stack is empty")
		return
	}

	for current := s.Head; current != nil; current = current.Next {
		fmt.Printf("%d ", current.Value)
	}

	fmt.Println()
}
