package main

import "fmt"

func main()  {
	res := twoSumOne([]int{1, 3, 2, 4, 3}, 4)
	fmt.Println(res)
}

func twoSumOne(s []int , sum int) []int {
	m := make(map[int]int)
	for i, v := range s {
		r, res := m[v]
		//值存在于map，则说明找到了差值
		if  res {
			fmt.Println(r, res, v, m)
			return []int{ r, i}
		}
		fmt.Println(r, res, v, m)
		//保存差值对应的索引
		m[sum - v] = i
	}
	return []int{}
}



