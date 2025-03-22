package arraydeque

import "fmt"

type ArrayDeque struct {
	size     int
	front    int
	capacity int
	a        [10]int
}

func NewArrayDeque() *ArrayDeque {
	return &ArrayDeque{
		size:     0,
		front:    0,
		capacity: 10,
		a:        [10]int{},
	}
}

func (d *ArrayDeque) IsEmpty() bool {
	return d.size == 0
}

func (d *ArrayDeque) IsFull() bool {
	return d.size == d.capacity
}

func (d *ArrayDeque) InsertFront(value int) {
	if d.IsFull() {
		fmt.Println("Deque full")
		return
	}

	d.front = (d.front - 1 + d.capacity) % d.capacity

	d.a[d.front] = value
	d.size++
}

func (d *ArrayDeque) DeleteFront() int {
	if d.IsEmpty() {
		fmt.Println("Deque empty")
		return 0
	}

	retVal := d.a[d.front]

	d.front = (d.front + 1) % d.capacity

	d.size--

	return retVal
}

func (d *ArrayDeque) InsertBack(value int) {
	if d.IsFull() {
		fmt.Println("Deque full")
		return
	}

	back := (d.front + d.size) % d.capacity

	d.a[back] = value
	d.size++
}

func (d *ArrayDeque) DeleteBack() int {
	if d.IsEmpty() {
		fmt.Println("Deque empty")
		return 0
	}

	back := (d.front + d.size - 1) % d.capacity

	val := d.a[back]
	d.a[back] = 0
	d.size--

	return val
}

func (d *ArrayDeque) FrontEle() int {
	if d.IsEmpty() {
		fmt.Println("Deque empty")
		return 0 // Or some error indication
	}
	return d.a[d.front]
}

func (d *ArrayDeque) BackEle() int {
	if d.IsEmpty() {
		fmt.Println("Deque empty")
		return 0 // Or some error indication
	}
	return d.a[(d.front+d.size-1)%d.capacity]
}

func (d *ArrayDeque) Print() {
	if d.IsEmpty() {
		fmt.Println("Deque is empty")
		return
	}
	for _, val := range d.a {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
}
