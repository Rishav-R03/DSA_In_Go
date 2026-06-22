package main

type Node struct {
	Data int
	Next *Node
}

func mergelist(l1 *Node, l2 *Node) *Node {
	dummy := &Node{}
	cur := dummy

	for l1 != nil && l2 != nil {
		if l1.Data < l2.Data {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}

		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}
	return dummy.Next
}
