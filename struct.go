package main

import "fmt"

type Vertex struct {
	X, Y int
	s    string
}

func changeVertex(v Vertex) {
	v.X = 1000
}
func changeVertex2(v *Vertex) {
	v.X = 1000
}

func main() {
	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)
	v.X = 100
	fmt.Println(v.X, v.Y)

	v2 := Vertex{X: 1}
	fmt.Println(v2)

	v3 := Vertex{1, 2, "test"}
	fmt.Println(v3)

	v4 := Vertex{}
	fmt.Println(v4)

	var v5 Vertex
	fmt.Println(v5)
	fmt.Printf("%T\n", v5)

	v6 := new(Vertex)
	fmt.Println(v6)
	fmt.Printf("%T\n", v6)

	v7 := &Vertex{}
	fmt.Printf("%T\n", v7)
}
