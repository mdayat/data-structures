package main

import (
	"errors"
	"fmt"
	"os"
)

type queue[T any] []T

func (q *queue[T]) push(item T) {
	*q = append(*q, item)
}

func (q *queue[T]) pop() (poppedItem T, err error) {
	if len(*q) > 0 {
		poppedItem = (*q)[0]
		*q = (*q)[1:]
		return poppedItem, nil
	}

	err = errors.New("queue is empty")
	return poppedItem, err
}

func (q *queue[T]) peek() (peepedItem T, err error) {
	if len(*q) > 0 {
		peepedItem = (*q)[0]
		return peepedItem, nil
	}

	err = errors.New("queue is empty")
	return peepedItem, err
}

func (q *queue[T]) size() int {
	return len(*q)
}

func main() {
	queueItems := queue[int]{1, 2, 3}
	fmt.Println("Original: ", queueItems)

	// 1. Push
	queueItems.push(2)
	fmt.Println("After Pushed: ", queueItems)

	// 2. Pop
	poppedItem, err := queueItems.pop()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Popped Item: ", poppedItem)
	fmt.Println("After Popped: ", queueItems)

	// 3. Peek
	peepedItem, err := queueItems.peek()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Peeped: ", peepedItem)

	// 4. Size
	fmt.Println("Size: ", queueItems.size())
}
