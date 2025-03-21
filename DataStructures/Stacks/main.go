package main

import (
	"fmt"
	"practice/stack/stackarray"
	"practice/stack/stacklist"
)

func main() {
	TestStackList()
}

func TestStackList() {
	stack := stacklist.NewStack()
	stack.Display()
	fmt.Println("Adding item to stack...")
	stack.Push(1)

	fmt.Printf("Stack: ")
	stack.Display()

	fmt.Println("Adding item to stack...")
	stack.Push(2)

	fmt.Println("Adding item to stack...")
	stack.Push(3)

	fmt.Println("Adding item to stack...")
	stack.Push(4)

	fmt.Println("Adding item to stack...")
	stack.Push(5)

	fmt.Printf("Stack: ")
	stack.Display()

	peeked := stack.Peek()
	fmt.Printf("Peeked: %d\n", peeked)

	fmt.Printf("Stack: ")
	stack.Display()

	popped := stack.Pop()

	fmt.Printf("Popped: %d\n", popped)
	fmt.Printf("Stack: ")
	stack.Display()
}

func TestStackArray() {
	stack := stackarray.NewFixedStack(10)
	fmt.Printf("Newly created stack: %v\n", stack)
	fmt.Printf("Is stack empty: %t\n", stack.IsEmpty())
	_ = stack.Push(10)
	_ = stack.Push(20)
	_ = stack.Push(30)
	_ = stack.Push(40)
	_ = stack.Push(50)
	_ = stack.Push(60)
	_ = stack.Push(70)
	_ = stack.Push(80)
	_ = stack.Push(90)
	_ = stack.Push(10)
	fmt.Printf("Stack: %v\n", stack)
	peeked := stack.Peek()
	fmt.Printf("Peeked Value: %d\n", peeked)
	fmt.Printf("Stack: %v\n", stack)
	fmt.Printf("Is stack empty: %t\n", stack.IsEmpty())
	fmt.Printf("Is stack full: %t\n", stack.IsFull())
	popped := stack.Pop()
	fmt.Printf("Popped value: %d\n", popped)
	fmt.Printf("Stack: %v\n", stack)

	peeked = stack.Peek()
	fmt.Printf("Peeked Value: %d\n", peeked)

	// stackarray.DynamicStackExample()
}
