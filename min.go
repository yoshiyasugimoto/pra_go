package main

import "fmt"

func main() {
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	//min := l[0]
	//for _, v := range l {
	//	if v < min {
	//		min = v
	//	}
	//}
	//fmt.Println(min)
	var min int

	for i, num := range l {
		if i == 0 {
			min = num
			continue
		}
		if min >= num {
			min = num
		}
	}
	fmt.Println(min)
}
