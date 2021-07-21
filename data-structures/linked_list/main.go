package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepand(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l *linkedList) deleteWithValue(v int) {
	previousToDelete := l.head
	for previousToDelete.next.data != v {
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func (l linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Println(toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
}

func main() {
	myList := linkedList{}
	node1 := &node{data: 7}
	node2 := &node{data: 6}
	node3 := &node{data: 5}
	node4 := &node{data: 4}
	node5 := &node{data: 3}
	node6 := &node{data: 2}
	node7 := &node{data: 1}

	myList.prepand(node1)
	myList.prepand(node2)
	myList.prepand(node3)
	myList.prepand(node4)
	myList.prepand(node5)
	myList.prepand(node6)
	myList.prepand(node7)

	myList.printListData()
	myList.deleteWithValue(5)
	myList.printListData()
}
