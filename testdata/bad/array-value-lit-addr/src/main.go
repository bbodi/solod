package main

type Grid struct {
	cells [2]int
}

func main() {
	cells := [2]int{1, 2}
	// The literal is a statement expression, which has no address.
	g := &Grid{cells: cells}
	_ = g
}
