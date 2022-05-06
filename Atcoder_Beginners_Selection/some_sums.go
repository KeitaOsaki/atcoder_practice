package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	var splitC []int
	var list []int
	ans := 0

	fmt.Scan(&a, &b, &c)

	for i := 0; i <= a; i++ {
		if i > 9 {
			listC := splitInt(i, splitC)
			sum := 0
			for _, v := range listC {
				sum += v
			}
			if b <= sum && sum <= c {
				list = append(list, i)
			}
		} else {
			if b <= i && i <= c {
				list = append(list, i)

			}
		}
	}
	for _, v := range list {
		ans += v
	}
	fmt.Println(ans)
}

func splitInt(i int, list []int) []int {
	if i > 0 {
		return splitInt(i/10, append(list, i%10))
	}
	return list
}
