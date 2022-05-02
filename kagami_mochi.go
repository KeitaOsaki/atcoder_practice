package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	ans := 1

	fmt.Scan(&n)
	nums := scans(n)
	sort.Ints(nums)
	beforeLen := nums[0]
	for _, v := range nums {
		if beforeLen < v {
			beforeLen = v
			ans++
		}
	}
	fmt.Println(ans)
}

func scans(len int) (nums []int) {
	var num int
	for i := 0; i < len; i++ {
		fmt.Scanf("%d", &num)
		nums = append(nums, num)
	}
	return
}
