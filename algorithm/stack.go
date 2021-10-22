package main

import (
	"fmt"
	"sync"
)

/**
 * 数组栈
 */
type ArrayStack struct {
	array[] string
	size int
	lock sync.Mutex
}

//入栈
func (stack * ArrayStack) Push (v string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	stack.array = append(stack.array, v)
	stack.size = stack.size + 1
}

//出栈
func (stack * ArrayStack) Pop() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.size == 0 {
		panic("empty")
	}

	v := stack.array[stack.size - 1]
	newArray := make([]string, stack.size - 1, stack.size - 1)
	for i := 0; i < stack.size - 1; i++ {
		newArray[i] = stack.array[i]
	}

	stack.array = newArray
	stack.size = stack.size - 1
	return  v
}

//栈的大小
func (stack * ArrayStack) Len() int{
	return stack.size
}

//判断是否为空栈
func (stack * ArrayStack)isEmpty() bool {
	return stack.size == 0
}
//栈顶元素
func (stack * ArrayStack) Top() string {
	if stack.size == 0 {
		panic("empty")
	}

	return stack.array[stack.size - 1]
}

//链表栈
type LinkStack struct {
	root *LinkNode
	size int
	lock sync.Mutex
}

type LinkNode struct {
	Next *LinkNode
	Value string
}

func (stack * LinkStack) Push(v string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.root == nil {
		stack.root = new(LinkNode)
		stack.root.Value = v
	} else {
		//否则新元素插入链表的头部
		//原来的链表
		preNode := stack.root

		// 新节点
		newNode := new (LinkNode)
		newNode.Value = v
		newNode.Next = preNode
		stack.root = newNode
	}

	stack.size = stack.size + 1
}

func (stack * LinkStack) Pop() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.size == 0 {
		panic("empty")
	}

	//顶部元素出栈
	topNode := stack.root
	v := topNode.Value

	//将顶部元素的后继元素连接上
	stack.root = topNode.Next

	stack.size = stack.size - 1

	return v
}

func (stack * LinkStack) Top() string  {
	if stack.size == 0 {
		panic("empty")
	}

	return stack.root.Value
}

func (stack *LinkStack) Len() int {
	return stack.size
}

func main() {
   /*arrayStack := ArrayStack{}
   arrayStack.Push("a")
   fmt.Println(arrayStack.array)
   arrayStack.Push("b")
   fmt.Println(arrayStack.array)
   fmt.Println(arrayStack.Len())
   fmt.Println(arrayStack.Top())
   arrayStack.Pop()
   fmt.Println(arrayStack.array)*/

	linkStack := new(LinkStack)
	linkStack.Push("cat")
	linkStack.Push("dog")
	linkStack.Push("snack")

	fmt.Println("size :", linkStack.Len())
	fmt.Println("pop:", linkStack.Pop())
	fmt.Println("size:", linkStack.Len())
	linkStack.Push("cow")
	fmt.Println("pop:", linkStack.Pop())
}

