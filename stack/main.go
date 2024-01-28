package main

import (
	"errors"
	"fmt"
	"os"
)

type stack[T any] []T

func (s *stack[T]) push(item T) {
	*s = append(*s, item)
}

func (s *stack[T]) pop() (poppedItem T, err error) {
	if len((*s)) > 0 {
		poppedItem = (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return poppedItem, nil
	}

	err = errors.New("stack is empty")
	return poppedItem, err
}

func (s *stack[T]) peek() (peepedItem T, err error) {
	if len(*s) > 0 {
		peepedItem = (*s)[len(*s)-1]
		return peepedItem, nil
	}

	err = errors.New("stack is empty")
	return peepedItem, err
}

func (s *stack[T]) size() int {
	return len(*s)
}

func main() {
	stackItems := stack[int]{1, 2, 3}
	fmt.Println("Original: ", stackItems)

	// 1. Push
	stackItems.push(2)
	fmt.Println("Push: ", stackItems)

	// 2. Pop
	poppedItem, err := stackItems.pop()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Popped Value: ", poppedItem)
	fmt.Println("Pop: ", stackItems)

	// 3. Peek
	peepedItem, err := stackItems.peek()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Peek: ", peepedItem)

	// 4. Size
	fmt.Println("Size: ", stackItems.size())
}
