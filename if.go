package main

import "fmt"

func main() {
	var num int = 6
	if num%2 == 0 {
		fmt.Println("by 2")
	} else if num%3 == 0 {
		fmt.Println("by 3")
	} else {
		fmt.Println("else")
	}

	x, y := 11, 12
	if x == 10 && y == 10 {
		fmt.Println("&&")
	}
	if x == 11 || y == 12 {
		fmt.Println("||")
	}
}
