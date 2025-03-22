package main

import (
	"fmt"
	"practice/queue/circularqueue"
	"practice/queue/listqueue"
)

func main() {
	TestListQueue()
}

func TestListQueue() {
	queue := listqueue.NewQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Display()

	fmt.Printf("Dequeued: %d\n", queue.Dequeue())
	fmt.Printf("Dequeued: %d\n", queue.Dequeue())
	fmt.Printf("Dequeued: %d\n", queue.Dequeue())
	queue.Display()

	queue.Enqueue(3)
	queue.Enqueue(7)
	queue.Enqueue(9)
	queue.Display()
}

func TestCircularArrayQueue() {
	queue := circularqueue.NewQueue(5)
	queue.Display()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)
	queue.Display()

	frontEle := queue.GetFront()
	fmt.Printf("Front element is: %d\n", frontEle)
	remEle := queue.Dequeue()
	fmt.Printf("Removed; %d\n", remEle)
	queue.Display()
	queue.Enqueue(9000)
	queue.Display()
	frontEle = queue.GetFront()
	fmt.Printf("Front element is: %d\n", frontEle)
}
