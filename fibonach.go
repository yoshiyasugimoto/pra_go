package main

import "fmt"

func Fibo(n int) int {
	if n < 2 {
		return n
	}
	return Fibo(n-2) + Fibo(n-1)
}

func main() {
	n := 10
	fmt.Println(Fibo(n))
}
