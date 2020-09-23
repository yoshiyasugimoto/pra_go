package main

import "fmt"

func Fibo(n int) int {
	if n < 2 {
		return n
	}
	return Fibo(n-2) + Fibo(n-1)
}

func Fibo2(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func main() {
	n := 10
	fmt.Println(Fibo(n))
	fmt.Println(Fibo2(n))
}
