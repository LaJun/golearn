package main

import (
	"container/list"
	"fmt"
)

type CQueue struct {
	stack1, stack2 *list.List
}

func Constructor() *CQueue{
	return &CQueue{
		stack1: list.New(),
		stack2: list.New(),
	}
}

func (cq *CQueue) AppendTail(value int) {
	cq.stack1.PushBack(value)
}

func (cq *CQueue) DeleteHead() int {
	if cq.stack2.Len() == 0 {
		for cq.stack1.Len() > 0 {
			cq.stack2.PushBack(cq.stack1.Remove(cq.stack1.Back()))
		}
	}
	if cq.stack2.Len() != 0 {
		e := cq.stack2.Back()
		cq.stack2.Remove(e)
		return e.Value.(int)
	}
	return -1
}

func main() {
	Constructor().AppendTail(3)
	Constructor().AppendTail(2)
	Constructor().AppendTail(4)
	fmt.Println(Constructor().DeleteHead())
	Constructor().AppendTail(5)
	fmt.Println(Constructor().DeleteHead())
}