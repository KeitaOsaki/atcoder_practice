package main

import "fmt"

func main() {
	var num int
	var list []int
	ans := 0
	fmt.Scanf("%d", &num)

	listNum := splitInt(num, list)

	for _, v := range listNum {
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
