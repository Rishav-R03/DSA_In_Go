package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func insertAtHead(head *Node, val int) *Node {
	newNode := &Node{Data: val}
	if head == nil {
		return newNode
	}
	newNode.Next = head
	return newNode
}

func createList(slice []int) *Node {
	if len(slice) <= 0 {
		return nil
	}
	head := &Node{Data: slice[0]}

	curr := head
	for i := 1; i < len(slice); i++ {
		newNode := &Node{Data: slice[i]}
		curr.Next = newNode
		curr = curr.Next
	}
	return head
}

func printList(head *Node) {
	if head == nil {
		return
	}
	curr := head
	for curr != nil {
		fmt.Printf("%d ->", curr.Data)
		curr = curr.Next
	}
	fmt.Printf("null\n")
}

func createCycle(nums []int, cycleInd int) *Node {
	if len(nums) == 0 {
		return nil
	}
	head := &Node{Data: nums[0]}
	cur := head
	var cycleNode *Node
	if cycleInd == 0 {
		cycleNode = head
	}
	for i := 1; i < len(nums); i++ {
		cur.Next = &Node{Data: nums[i]}
		cur = cur.Next
		if i == cycleInd {
			cycleNode = cur
		}

		if cycleNode != nil {
			cur.Next = cycleNode
		}
	}
	return head
}
func detectCycle(head *Node) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}
	return false
}

func main() {
	slice := []int{10, 20, 30, 40, 50}
	ll := createList(slice)
	// fmt.Println(ll)
	cl := createCycle(slice, 3)
	printList(ll)
	if detectCycle(ll) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	if detectCycle(cl) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
