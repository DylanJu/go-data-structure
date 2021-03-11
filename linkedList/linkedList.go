package linkedlist

import "fmt"

type node struct {
	value int
	next  *node
}

// LinkedList is a kind of data structure made of a collection of node sequentially
type LinkedList struct {
	head *node
	tail *node
}

func (l *LinkedList) addNode(val int) {
	curr := l.head

	for curr.next != nil {
		curr = curr.next
	}

	curr.next = &node{value: val}
	l.tail = curr.next
}

func (l *LinkedList) prependNode(val int) {
	head := &node{value: val, next: l.head}
	l.head = head
}

func (l *LinkedList) insertNode(val int, target int) {
	curr := l.head

	for curr.value != target {
		curr = curr.next
	}

	newNode := &node{value: val, next: curr.next}
	curr.next = newNode

	if l.tail == curr {
		l.tail = newNode
	}
}

func (l *LinkedList) removeNode(val int) {
	if l.head.value == val {
		l.head = l.head.next
		return
	}

	prev := l.head
	curr := prev.next

	for curr.next != nil {
		if curr.value == val {
			prev.next = curr.next
			break
		} else {
			prev = curr
			curr = curr.next
		}
	}

	if l.tail.value == val {
		prev.next = nil
		l.tail = prev
	}
}

func (l *LinkedList) printNodes() {
	curr := l.head

	for curr.next != nil {
		fmt.Print(curr.value, " --> ")
		curr = curr.next
	}
	fmt.Println(curr.value)
}

// MakeLinkedList makes linked-list
func MakeLinkedList(val int) LinkedList {
	head := &node{value: val}
	LinkedList := LinkedList{head: head}
	return LinkedList
}
