package listdeque

import "fmt"

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type ListDeque struct {
	front *Node
	back  *Node
}

func NewNode(value int) *Node {
	return &Node{
		value: value,
		next:  nil,
		prev:  nil,
	}
}

func NewListDeque() *ListDeque {
	return &ListDeque{
		front: nil,
		back:  nil,
	}
}

func (d *ListDeque) InsertFront(value int) {
	newNode := NewNode(value)
	if d.front == nil {
		d.front = newNode
		d.back = d.front
		return
	}
	newNode.next = d.front
	d.front.prev = newNode
	d.front = newNode
}

func (d *ListDeque) InsertBack(value int) {
	newNode := NewNode(value)
	if d.back == nil {
		d.front = newNode
		d.back = d.front
		return
	}
	newNode.prev = d.back
	d.back.next = newNode
	d.back = newNode
}

func (d *ListDeque) DeleteFront() {
	if d.front == nil {
		fmt.Println("Underflow")
		return
	}

	temp := d.front
	d.front = d.front.next
	d.front.prev = nil

	temp.next = nil
}

func (d *ListDeque) DeleteBack() {
	if d.back == nil {
		fmt.Println("Underflow")
		return
	}

	temp := d.back
	d.back = d.back.prev
	d.back.next = nil

	temp.prev = nil
}

func (d *ListDeque) FrontEle() int {
	if d.front == nil {
		fmt.Println("Deque empty")
		return 0
	}
	return d.front.value
}

func (d *ListDeque) BackEle() int {
	if d.back == nil {
		fmt.Println("Deque empty")
		return 0
	}
	return d.back.value
}

func (d *ListDeque) PrintDeque() {
	if d.front == nil {
		fmt.Println("Deque empty")
		return
	}
	fmt.Printf("Deque: ")

	for current := d.front; current != nil; current = current.next {
		fmt.Printf("%d ", current.value)
	}

	fmt.Println()
}
