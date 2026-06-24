package main

import (
	"errors"
	"fmt"
	"log"
)

type Node struct {
	data int
	next *Node
}

type Stack struct {
	top  *Node
	size int
}

// Create stack with linkedlist
func (s *Stack) Push(val int) {
	newNode := &Node{data: val, next: s.top}
	s.top = newNode
	s.size++
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack underflow: cannot pop")
	}
	popped := s.top.data
	s.top = s.top.next
	s.size--
	return popped, nil
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack) Peek() int {
	return s.top.data
}

func (s *Stack) Size() int {
	return s.size
}

func main() {
	stack := &Stack{}
	fmt.Printf("Pushing elements: 10, 20, 30\n")
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Printf("current size of stack: %d \n", stack.size)
	top := stack.Peek()
	fmt.Printf("current top val: %d\n", top)

	fmt.Printf("popping elements")
	popped, err := stack.Pop()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("popped value: %d\n", popped)

}
