package main

type Node struct {
	Data int
	Next *Node
}

func GetMiddleNode(head *Node) *Node {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
