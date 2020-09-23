package main

import "fmt"

type Ver struct {
	X, Y int
	s    string
}

func changeVertex(v Ver) {
	v.X = 1000
}
func changeVertex2(v Ver) {
	v.X = 1000
}

func main() {
	v := Ver{X: 1, Y: 2}
	fmt.Println(v)
	v.X = 100
	fmt.Println(v.X, v.Y)

	v2 := Ver{X: 1}
	fmt.Println(v2)

	v3 := Ver{1, 2, "test"}
	fmt.Println(v3)

	v4 := Ver{}
	fmt.Println(v4)

	var v5 Ver
	fmt.Println(v5)
	fmt.Printf("%T\n", v5)

	v6 := new(Ver)
	fmt.Println(v6)
	fmt.Printf("%T\n", v6)

	v7 := &Ver{}
	fmt.Printf("%T\n", v7)
}
