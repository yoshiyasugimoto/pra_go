package main

import "fmt"

func foo(params ...int) {
	fmt.Println(len(params), params)
	for _, params := range params {
		fmt.Println(params)
	}
}

func main() {
	foo()
	foo(10, 20)
	foo(20, 30, 30)

	s := []int{1, 2, 3}
	fmt.Println(s)

	foo(s...)

}
