package main

import (
	"fmt"
	"log"
)

type Node struct {
	data int
	next *Node
}

// to make operations O(1)
type Queue struct {
	front *Node
	rear  *Node
}

func NewNode(data int, next *Node) *Node {
	return &Node{data: data, next: next}
}

func (q *Queue) Enqueue(data int) {
	newNode := NewNode(data, nil)
	if q.rear == nil {
		q.front = newNode
		q.rear = newNode
		return
	}

	q.rear.next = newNode
	q.rear = newNode
}

func (q *Queue) Dequeue() (int, bool) {
	if q.front == nil {
		log.Println("stack underflow cannot dequeue")
		return -1, false
	}
	val := q.front.data
	q.front = q.front.next
	if q.front == nil {
		q.rear = nil
	}
	return val, true
}

func (q *Queue) Peek() (int, bool) {
	if q.front == nil {
		return -1, false
	}
	return q.front.data, true
}

func (q *Queue) IsEmpty() bool {
	return q.front == nil
}

func main() {
	var q Queue
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	val, ok := q.Peek()
	if !ok {
		fmt.Println("Empty queue")
	}
	fmt.Println(val)
	q.Dequeue()
	q.Peek()
	val1, ok := q.Peek()
	if !ok {
		fmt.Println("Empty queue")
	}
	fmt.Println(val1)
}
