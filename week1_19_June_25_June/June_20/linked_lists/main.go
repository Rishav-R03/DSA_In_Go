package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func createList(slice []int) *ListNode {
	if len(slice) == 0 {
		return nil
	}
	head := &ListNode{Val: slice[0]}

	curr := head
	for i := 1; i < len(slice); i++ {
		newNode := &ListNode{Val: slice[i]}
		curr.Next = newNode
		curr = newNode
	}
	return head
}

func PrintList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Printf("%d -> ", cur.Val)
		cur = cur.Next
	}
	fmt.Println("nil")
}

func insertAtHead(head *ListNode, val int) *ListNode {
	newNode := &ListNode{Val: val}
	newNode.Next = head
	return newNode
}

func insertAtEnd(head *ListNode, val int) *ListNode {
	newNode := &ListNode{Val: val}
	if head == nil {
		return newNode
	}
	curr := head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = newNode
	return head
}

func deleteHead(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	return head.Next
}

func deleteEnd(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return nil
	}
	curr := head
	for curr.Next.Next != nil {
		curr = curr.Next
	}
	curr.Next = nil

	return head
}
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func main() {
	head := createList([]int{10, 20, 30})

	PrintList(head)

	head = insertAtHead(head, 5)
	PrintList(head)

	head = insertAtEnd(head, 40)
	PrintList(head)

	head = deleteHead(head)
	PrintList(head)

	head = deleteEnd(head)
	PrintList(head)

	head = reverseList(head)
	PrintList(head)
}
