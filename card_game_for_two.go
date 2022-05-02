package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	allcePoint := 0
	bobPoint := 0

	fmt.Scan(&n)
	nums := scans(n)
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	for i, v := range nums {
		if i%2 == 0 {
			allcePoint += v
		} else {
			bobPoint += v
		}
	}
	fmt.Println(allcePoint - bobPoint)

}

//func scans(len int) (nums []int) {
//	var num int
//	for i := 0; i < len; i++ {
//		fmt.Scanf("%d", &num)
//		nums = append(nums, num)
//	}
//	return
//}
