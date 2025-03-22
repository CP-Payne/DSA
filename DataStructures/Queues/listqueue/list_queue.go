package listqueue

import "fmt"

type Node struct {
	value int
	next  *Node
}

func NewNode(value int) *Node {
	return &Node{
		value: value,
		next:  nil,
	}
}

type Queue struct {
	front *Node
	rear  *Node
}

func NewQueue() *Queue {
	return &Queue{
		front: nil,
		rear:  nil,
	}
}

func (q *Queue) Display() {
	if q.front == nil {
		fmt.Println("Queue is empty")
		return
	}
	fmt.Printf("Queue: ")

	for current := q.front; current != nil; current = current.next {
		fmt.Printf("%d ", current.value)
	}
	fmt.Println()
}

func (q *Queue) Enqueue(val int) {
	newNode := NewNode(val)
	if q.rear == nil {
		q.rear = newNode
		q.front = newNode
		return
	}
	q.rear.next = newNode
	q.rear = newNode
}

func (q *Queue) Dequeue() int {
	if q.front == nil {
		fmt.Println("Queue is empty")
		return -1
	}
	ret := q.front.value
	q.front = q.front.next
	return ret
}
