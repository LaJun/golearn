package main

import (
	"fmt"
	"sync"
)

type  Array struct {
	array []int
	len int
	cap int
	lock sync.Mutex
}

func Make(len, cap int)* Array  {
	s := new(Array)
	if len > cap {
		panic("len large than cap")
	}

	array := make([]int, cap,cap)
	s.array = array
	s.cap = cap
	s.len = 0
	return s
}

func (a * Array) Append(element int) {
	a.lock.Lock()
	defer  a.lock.Unlock()

	if a.len == a.cap {
		newCap := 2 * a.len
		if a.cap == 0 {
			newCap = 1
		}
		newArray := make([]int,newCap, newCap)
		for k, v := range a.array {
			newArray[k] = v
		}
		a.array = newArray
		a.cap = newCap
	}
	a.array[a.len] = element
	a.len = a.len + 1
}

func (a *Array) Appends(elements ...int) {
	for _,v := range elements {
		a.Append(v)
	}
}

func (a * Array) Get(index int) int {
	if a.len == 0 || index >= a.len {
		panic("index over len")
	}
	return a.array[index]
}

func (a * Array) Len() int{
	return a.len
}

func (a *Array) Cap() int {
	return a.cap
}

func Print(array * Array)(result string) {
	result = "["
	for i := 0; i < array.Len(); i++ {
		// 第一个元素
		if i == 0 {
			result = fmt.Sprintf("%s%d", result, array.Get(i))
			continue
		}

		result = fmt.Sprintf("%s %d", result, array.Get(i))
	}
	result = result + "]"
	return
}

func main()  {
	a := Make(0, 3)
	fmt.Println("cap", a.Cap(), "len", a.Len(),"array:", Print(a) )

	// 增加一个元素
	a.Append(10)
	fmt.Println("cap", a.Cap(), "len", a.Len(), "array:", Print(a))

	// 增加一个元素
	a.Appends(9, 10, 30)
	fmt.Println("cap", a.Cap(), "len", a.Len(), "array:", Print(a))

	fmt.Println(a.Get(2))
}