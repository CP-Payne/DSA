package circularqueue

import "fmt"

type Queue struct {
	arr      []int
	capacity int
	size     int
	front    int
}

func NewQueue(size int) *Queue {
	return &Queue{
		arr:      make([]int, size),
		capacity: size,
		size:     0,
		front:    0,
	}
}

func (q *Queue) GetFront() int {
	if q.size == 0 {
		fmt.Println("Queue empty")
		return -1
	}
	return q.arr[q.front]
}

func (q *Queue) GetRear() int {
	if q.size == 0 {
		fmt.Println("Queue empty")
		return -1
	}
	rear := (q.front + q.size - 1) % q.capacity

	return q.arr[rear]
}

func (q *Queue) Enqueue(x int) {
	if q.size == q.capacity {
		fmt.Println("Queue is full")
		return
	}
	rear := (q.front + q.size) % q.capacity

	q.arr[rear] = x
	q.size++
}

func (q *Queue) Dequeue() int {
	if q.size == 0 {
		fmt.Println("Queue is empty")
		return -1
	}
	res := q.arr[q.front]

	q.front = (q.front + 1) % q.capacity
	q.size--
	return res
}

func (q *Queue) Display() {
	if q.size == 0 {
		fmt.Println("Queue is empty")
		return
	}

	fmt.Printf("Queue: %v\n", q.arr)
}
