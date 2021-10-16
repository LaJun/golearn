package main

import (
	"fmt"
	"sort"
)

func main()  {
	s := []int{13, 12, 12, 13, 12, 34}
	//num := wayOne(s)
	num := wayTwo(s)
	fmt.Println(num)
}

func wayOne(s []int) int {
	m := make(map[int]int)
	for _,v:= range s {
		if _, ok := m[v]; ok {
			return v
		}
		m[v] = 1
	}
	return -1
}

func wayTwo(s []int) int{
	sort.Ints(s)
	for i, numSize := 0, len(s); i < numSize-1; i++ {
		if s[i] == s[i+1] {
			return s[i]
		}
	}
	return -1
}
