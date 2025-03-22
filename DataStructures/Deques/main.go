package main

import (
	"fmt"
	"practice/deque/arraydeque"
	"practice/deque/listdeque"
)

func main() {
	listDeque := listdeque.NewListDeque()
	listDeque.PrintDeque()
	listDeque.InsertFront(1)
	listDeque.InsertFront(2)
	listDeque.InsertFront(3)
	listDeque.PrintDeque()
	fmt.Printf("Front element: %d\n", listDeque.FrontEle())
	fmt.Printf("Back element: %d\n", listDeque.BackEle())
	listDeque.InsertBack(9)
	listDeque.InsertBack(8)
	listDeque.InsertBack(7)
	listDeque.PrintDeque()
	fmt.Printf("Front element: %d\n", listDeque.FrontEle())
	fmt.Printf("Back element: %d\n", listDeque.BackEle())
}

func TestArrayDeque() {
	arrayDeque := arraydeque.NewArrayDeque()
	arrayDeque.InsertFront(10)
	arrayDeque.Print()
	arrayDeque.InsertFront(20)
	arrayDeque.Print()

	arrayDeque.InsertBack(30)
	arrayDeque.Print()
	arrayDeque.InsertBack(40)
	arrayDeque.Print()

	arrayDeque.InsertFront(30)
	arrayDeque.Print()
	arrayDeque.InsertFront(35)
	arrayDeque.Print()

	arrayDeque.InsertBack(55)
	arrayDeque.Print()
	arrayDeque.InsertBack(53)
	arrayDeque.Print()

	arrayDeque.InsertBack(65)
	arrayDeque.Print()
	arrayDeque.InsertBack(58)
	arrayDeque.Print()

	arrayDeque.InsertBack(58)
	arrayDeque.Print()

	val := arrayDeque.DeleteBack()
	fmt.Printf("Value Deleted From Back: %d\n", val)

	fmt.Println("Inserting new value in back")
	arrayDeque.InsertBack(1000)
	arrayDeque.Print()

	val = arrayDeque.DeleteFront()
	fmt.Printf("Value Deleted From Front: %d\n", val)

	fmt.Println("Inserting new value in Front")
	arrayDeque.InsertFront(2000)
	arrayDeque.Print()

	frontEle := arrayDeque.FrontEle()
	fmt.Printf("Front Element: %d\n", frontEle)
	backEle := arrayDeque.BackEle()
	fmt.Printf("Back Element: %d\n", backEle)
}
