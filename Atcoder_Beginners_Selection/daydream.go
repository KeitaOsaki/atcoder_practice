package main

import (
	"fmt"
	"strings"
)

func main() {

	var s string

	fmt.Scan(&s)
	s1 := s

	stringList := []string{"eraser", "erase", "dreamer", "dream"}

	for _, v := range stringList {
		//s1 = strings.TrimPrefix(s1, v)
		s1 = strings.Replace(s1, v, "", -1)
	}
	if s1 == "" {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}

//
//func main() {
//	var s string
//	var checkS []string
//
//	fmt.Scan(&s)
//
//	stringList := []string{"dream", "dreamer", "erase", "eraser"}
//	for _, v := range stringList {
//		checkS = append(checkS, v)
//		for _, n := range stringList {
//			if v != n {
//				checkS = append(checkS, v+n)
//			}
//			for _, j := range stringList {
//				if v != n && v != j && j != n {
//					checkS = append(checkS, v+n+j)
//				}
//				for _, k := range stringList {
//					if v != n && v != j && v != k && n != j && n != k && j != k {
//						checkS = append(checkS, v+n+j+k)
//					}
//				}
//			}
//		}
//	}
//	fmt.Println(len(checkS))
//
//	for _, v := range checkS {
//		if v == s {
//			fmt.Println("YES")
//			return
//		}
//	}
//	fmt.Println("NO")
//}
