package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int

	fmt.Scanf("%d", &n)

	nums := scans(n)

	var resList []int

	for i := 0; i < n; i++ {
		res := count(nums[i])
		resList = append(resList, res)
	}
	sort.Ints(resList)
	fmt.Println(resList[0])
	return
}

func count(n int) int {
	countNum := 0
	for n%2 == 0 {
		countNum++
		n = n / 2
	}
	return countNum
}

//func scans(len int) (nums []int) {
//	var num int
//	for i := 0; i < len; i++ {
//		fmt.Scanf("%d", &num)
//		nums = append(nums, num)
//	}
//	return
//}
