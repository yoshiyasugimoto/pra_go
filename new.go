package main

import "fmt"

func main() {
	var p *int = new(int)
	fmt.Println(p)
	*p++
	fmt.Println(*p)

	var p2 *int
	fmt.Println(p2)
	fmt.Printf("%T\n",p2)
}
