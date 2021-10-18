package main

import (
	"fmt"
	"sort"
)

//在有序二维数组中查找中查找一个数字
func main()  {
	matrix :=  [][]int{
		[]int{1, 4, 7, 11, 15},
		[]int{2, 5, 8, 12, 19},
		[]int{3, 6, 9, 16, 22},
		[]int{10, 13, 14, 17, 24},
		[]int{18, 21, 23, 26, 30},
	}
	i, j := findNumMatrixWayTwo(matrix, 24)
	fmt.Println(i, j)
}

//从最左下角开始， j++, i--,如果当前值大于target，j++，不然i--，相等则返回
func findNumMatrixWayOne(matrix [][]int, target int) (int, int)  {
	j := 0
	i := len(matrix) - 1
	for i > -1 {
		if j < len(matrix[i]) {
			if target < matrix[i][j] {
				i --
			} else if target > matrix[i][j] {
				j ++
			} else if target == matrix[i][j] {
				return i, j
			}
		} else  {
			return -1, -1
		}
	}
	return -1, -1
}

func findNumMatrixWayTwo(matrix [][]int, target int) (int, int)  {
	for i,nums := range matrix {
		//返回可以插入的位置
		j := sort.SearchInts(nums, target)
		if j < len(nums) && target == nums[j] {
			return i, j
		}
	}
	return -1, -1
}
