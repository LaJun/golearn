package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//反转链表的实现
func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode = nil
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre //这句话最重要
	}
	return pre
}

func main() {
	head := new(ListNode)
	head.Val = 1
	for i:= 2; i < 9; i++ {
		node := new(ListNode)
		node.Val = i
		head.Next = node
	}
	fmt.Println(head.Next)
	pre := reverseList(head)
	fmt.Println(pre.Next)
}