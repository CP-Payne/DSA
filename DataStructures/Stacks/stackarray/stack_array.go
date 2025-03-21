package stackarray

import "fmt"

type Stack struct {
	top int
	cap int
	a   [10]int
}

func NewFixedStack(cap int) *Stack {
	return &Stack{
		top: -1,
		cap: cap,
		a:   [10]int{},
	}
}

func (s *Stack) Push(val int) bool {
	if s.top == s.cap-1 {
		fmt.Println("Stack Overflow")
		return false
	}
	s.top += 1
	s.a[s.top] = val
	return true
}

func (s *Stack) Pop() int {
	if s.top == -1 {
		fmt.Println("Stack Underflow")
		return 0
	}

	retVal := s.a[s.top]
	s.top -= 1

	return retVal
}

func (s *Stack) Peek() int {
	if s.top == -1 {
		fmt.Println("Stack Underflow")
		return 0
	}

	return s.a[s.top]
}

func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

func (s *Stack) IsFull() bool {
	return s.top == s.cap-1
}

func DynamicStackExample() {
	stack := []int{}

	// Push Elements
	stack = append(stack, 10)
	stack = append(stack, 20)
	stack = append(stack, 30)

	fmt.Printf("Stack Items: %v\n", stack)

	// Pop and print the top element
	fmt.Printf("%d popped from stack\n", stack[len(stack)-1])
	stack = stack[:len(stack)-1]

	fmt.Printf("Top element is: %d\n", stack[len(stack)-1])

	fmt.Printf("Stack Items: %v\n", stack)
}
