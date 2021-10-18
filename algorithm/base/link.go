package main

import (
	"fmt"
)

type LinkNode struct {
	Data int64
	NextNode *LinkNode
}

func main()  {
	createLink()
	ring := createRingLink(4)
	fmt.Println(ring.prev.prev.prev.prev.prev)

}

func createLink(){
	node := new (LinkNode)
	node.Data = 2

	//新增节点
	node1 := new (LinkNode)
	node1.Data = 3
	node.NextNode = node1

	//新增节点
	node2 := new(LinkNode)
	node2.Data = 4
	node1.NextNode = node2

	nowNode := node
	for  {
		if nowNode != nil {
			fmt.Println(nowNode.Data)
			nowNode = nowNode.NextNode
			continue
		}
		break
	}
}

type Ring struct {
	next, prev * Ring
	Value interface{}
}

func createRingLink(n int) *Ring{
	if n <= 0 {
		return nil
	}
	//ring.Ring{}
	//初始化空的循环链表，前驱后驱都指向自己
	/*r := new(Ring)
	r.next = r
	r.prev = r*/
	r := new(Ring)
	p := r

	for i := 1; i < n; i++ {
		p.next = &Ring{prev:p}
		p = p.next
		p.Value = i
	}

	p.next = r
	r.prev = p
	return r
}
