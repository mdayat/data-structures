package main

import (
	"errors"
	"fmt"
	"os"
)

type node struct {
	data int
	next *node
}

type singlyLinkedList struct {
	head *node
	tail *node
}

func (l *singlyLinkedList) pushToHead(val int) {
	var node *node = &node{data: val}

	if l.head == nil {
		l.head, l.tail = node, node
	} else {
		node.next = l.head
		l.head = node
	}
}

func (l *singlyLinkedList) pushToTail(val int) {
	var node *node = &node{data: val}

	if l.head == nil {
		l.head, l.tail = node, node
	} else {
		l.tail.next = node
		l.tail = node
	}
}

func (l *singlyLinkedList) popFromHead() (poppedNodeVal int, err error) {
	if l.head == nil {
		err = errors.New("linked list is empty")
		return poppedNodeVal, err
	}

	if l.head == l.tail {
		poppedNodeVal = l.head.data
		l.head, l.tail = nil, nil
		return poppedNodeVal, nil
	}

	poppedNodeVal = l.head.data
	l.head = l.head.next
	return poppedNodeVal, nil
}

func (l *singlyLinkedList) popFromTail() (poppedNodeVal int, err error) {
	if l.head == nil {
		err = errors.New("linked list is empty")
		return poppedNodeVal, err
	}

	if l.head == l.tail {
		poppedNodeVal = l.head.data
		l.head, l.tail = nil, nil
		return poppedNodeVal, nil
	}

	currNode := l.head
	for currNode.next != nil && currNode.next != l.tail {
		currNode = currNode.next
	}

	poppedNodeVal = l.tail.data
	l.tail = currNode
	l.tail.next = nil
	return poppedNodeVal, nil
}

func (l *singlyLinkedList) searchAndPop(val int) (poppedNodeVal int, err error) {
	if l.head == nil {
		err = errors.New("linked list is empty")
		return poppedNodeVal, err
	}

	if l.head.data == val {
		poppedNodeVal = l.head.data

		if l.head.next != nil {
			l.head = l.head.next
		} else {
			l.head = nil
		}

		return poppedNodeVal, nil
	}

	if l.tail.data == val {
		poppedNodeVal = l.tail.data
		currNode := l.head

		for currNode.next != nil && currNode.next != l.tail {
			currNode = currNode.next
		}

		l.tail = currNode
		l.tail.next = nil
		return poppedNodeVal, nil
	}

	currNode := l.head
	for currNode.next != l.tail {
		if currNode.next.data == val {
			poppedNodeVal = currNode.next.data
			currNode.next = currNode.next.next
			return poppedNodeVal, err
		}
		currNode = currNode.next
	}

	err = errors.New("item not found")
	return poppedNodeVal, err
}

func (l *singlyLinkedList) searchAndPeek(val int) (peepedNodeVal int, err error) {
	if l.head == nil {
		err = errors.New("linked list is empty")
		return peepedNodeVal, err
	}

	if l.head.data == val {
		peepedNodeVal = l.head.data
		return peepedNodeVal, nil
	}

	if l.tail.data == val {
		peepedNodeVal = l.tail.data
		return peepedNodeVal, nil
	}

	currNode := l.head
	for currNode.next != l.tail {
		if currNode.next.data == val {
			peepedNodeVal = currNode.next.data
			return peepedNodeVal, err
		}
		currNode = currNode.next
	}

	err = errors.New("item not found")
	return peepedNodeVal, err
}

func main() {
	sll := new(singlyLinkedList)

	// Push To Head
	sll.pushToHead(2)
	sll.pushToHead(1)

	// Push To Tail
	sll.pushToTail(3)
	sll.pushToTail(4)

	// Search and Peek
	peepedVal, err := sll.searchAndPeek(3)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Search and Peek: ", peepedVal)

	// Search and Pop
	poppedVal, err := sll.searchAndPop(3)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Search and Pop: ", poppedVal)

	// Pop From Head
	headVal, err := sll.popFromHead()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Pop Head: ", headVal)

	// Pop From Tail
	tailVal, err := sll.popFromTail()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Pop Tail: ", tailVal)
}
